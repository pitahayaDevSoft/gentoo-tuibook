package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

//go:embed data/handbook
var handbookFS embed.FS

var (
	Version = "0.1.0"
)

const (
	handbookRoot     = "data/handbook"
	defaultFetchDate = "2026-05-23"
)

func getHandbookPath(parts ...string) string {
	return strings.Join(append([]string{handbookRoot}, parts...), "/")
}

var (
	editLinkRe     = regexp.MustCompile(`\\\[[ \t]*\[[ \t]*[Ee]dit[ \t]*\][ \t]*\([^)]*\)[ \t]*\\\]`)
	langLinkRe     = regexp.MustCompile(`(?m)^Other languages:\n\n(?:- .*\n)+\n`)
	markdownLinkRe = regexp.MustCompile(`\[([^\]]+)\]\(([^\s\)]+)(?:\s+"[^"]*")?\)`)
	rawURLRe       = regexp.MustCompile(`https?://[^\s\)"]+`)

	glamourStyleTemplate = `{
		"document": { "block_prefix": "\n", "block_suffix": "\n", "color": "#EBEBEB", "margin": 1 },
		"heading": { "block_suffix": "\n", "color": "{{brandPrimary}}", "bold": true },
		"h1": { "prefix": " # ", "suffix": " ", "color": "{{brandPrimary}}", "background_color": "#1c2130", "bold": true },
		"h2": { "prefix": " ## ", "color": "{{brandPrimary}}", "bold": true },
		"h3": { "prefix": " ### ", "color": "{{brandPrimary}}" },
		"strong": { "bold": true, "color": "#FFFFFF" },
		"emph": { "italic": true, "color": "#94A3B8" },
		"link": { "color": "{{brandPrimary}}", "underline": true },
		"link_text": { "color": "{{brandPrimary}}", "bold": true },
		"code": { "prefix": " ", "suffix": " ", "color": "#F0C674", "background_color": "#1C2130" },
		"code_block": { "margin": 1, "chroma": {
			"text": { "color": "#C4C4C4" },
			"keyword": { "color": "{{brandPrimary}}" },
			"name_function": { "color": "#7FDBCA" },
			"literal_string": { "color": "#F0C674" },
			"comment": { "color": "#6272A4" },
			"name_builtin": { "color": "#7FDBCA" },
			"operator": { "color": "#EF8080" },
			"literal_number": { "color": "#6EEFC0" },
			"keyword_type": { "color": "#B294BB" },
			"name_class": { "color": "#FFFFFF", "underline": true, "bold": true },
			"name_attribute": { "color": "#7A7AE6" },
			"name_tag": { "color": "#B083EA" }
		}},
		"item": { "block_prefix": "- ", "color": "#EBEBEB" },
		"enumeration": { "block_prefix": ". " },
		"hr": { "color": "{{brandPrimary}}" },
		"block_quote": { "color": "#6272A4", "indent": 1, "indent_token": "| " },
		"table": { "center_separator": "+", "column_separator": "|", "row_separator": "-" }
	}`

	glamourStyle []byte

	showWelcome     = true
	customStylePath = ""
	wrapWidth       = 0

	surfaceBg    = lipgloss.Color("#060608")
	surface3     = lipgloss.Color("#1c2130")
	brandPrimary = lipgloss.Color("#ff007f")

	baseStyle     = lipgloss.NewStyle().Background(surfaceBg).Foreground(lipgloss.Color("#EBEBEB"))
	statusStyle   = lipgloss.NewStyle().Background(brandPrimary).Foreground(surfaceBg).Padding(0, 2).Bold(true)
	selectedStyle = lipgloss.NewStyle().Foreground(brandPrimary).Bold(true)

	// Estilos segmentados de la barra de estado skeuomórfica
	modeBadgeStyle = lipgloss.NewStyle().Background(brandPrimary).Foreground(surfaceBg).Padding(0, 1).Bold(true)
	infoBadgeStyle = lipgloss.NewStyle().Background(surface3).Foreground(brandPrimary).Padding(0, 1).Bold(true)
	helpBadgeStyle = lipgloss.NewStyle().Background(surfaceBg).Foreground(lipgloss.Color("#888888")).Padding(0, 1)

	sidebarCfg = SidebarConfig{WidthPercent: 25, Min: 20, Max: 40}
)

// --- Config ---

type SidebarConfig struct {
	WidthPercent int `json:"width_percent"`
	Min          int `json:"min"`
	Max          int `json:"max"`
}

type ThemeColors struct {
	BrandPrimary string `json:"brand_primary"`
	SurfaceBg    string `json:"surface_bg"`
	Surface3     string `json:"surface_3"`
	Text         string `json:"text"`
}

type Config struct {
	DefaultArch     string        `json:"default_arch"`
	DefaultLang     string        `json:"default_lang"`
	ShowWelcome     bool          `json:"show_welcome"`
	CustomStylePath string        `json:"custom_style_path"`
	WrapWidth       int           `json:"wrap_width"`
	Sidebar         SidebarConfig `json:"sidebar"`
	Theme           ThemeColors   `json:"theme"`
}

func defaultConfig() Config {
	return Config{
		DefaultArch:     "auto",
		DefaultLang:     "auto",
		ShowWelcome:     true,
		CustomStylePath: "",
		WrapWidth:       0,
		Sidebar:         SidebarConfig{WidthPercent: 25, Min: 20, Max: 40},
		Theme: ThemeColors{
			BrandPrimary: "#ff007f",
			SurfaceBg:    "#060608",
			Surface3:     "#1c2130",
			Text:         "#EBEBEB",
		},
	}
}

func configDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		return filepath.Join(os.Getenv("HOME"), ".config", "gentoo-tuibook")
	}
	return filepath.Join(dir, "gentoo-tuibook")
}

func configPath() string {
	return filepath.Join(configDir(), "config.json")
}

func loadConfig() Config {
	cfg := defaultConfig()
	data, err := os.ReadFile(configPath())
	if err != nil {
		return cfg
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Printf("config parse: %v", err)
		return defaultConfig()
	}

	// Migrar el viejo color cian por defecto (#22d3ee) al nuevo rosa Gentoo (#ff007f)
	if cfg.Theme.BrandPrimary == "#22d3ee" {
		cfg.Theme.BrandPrimary = "#ff007f"
		updatedData, err := json.MarshalIndent(cfg, "", "  ")
		if err == nil {
			os.WriteFile(configPath(), updatedData, 0644)
		}
	}

	return cfg
}

func ensureConfig() {
	dir := configDir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("config dir: %v", err)
		return
	}
	if _, err := os.Stat(configPath()); err == nil {
		return
	}
	cfg := defaultConfig()
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Printf("config marshal: %v", err)
		return
	}
	if err := os.WriteFile(configPath(), data, 0644); err != nil {
		log.Printf("config write: %v", err)
	}
}

func applyConfig(cfg Config) {
	if cfg.Sidebar.WidthPercent > 0 {
		sidebarCfg.WidthPercent = cfg.Sidebar.WidthPercent
	}
	if cfg.Sidebar.Min > 0 {
		sidebarCfg.Min = cfg.Sidebar.Min
	}
	if cfg.Sidebar.Max > 0 {
		sidebarCfg.Max = cfg.Sidebar.Max
	}

	showWelcome = cfg.ShowWelcome
	customStylePath = cfg.CustomStylePath
	wrapWidth = cfg.WrapWidth

	textColor := lipgloss.Color("#EBEBEB")
	if cfg.Theme.BrandPrimary != "" {
		brandPrimary = lipgloss.Color(cfg.Theme.BrandPrimary)
	}
	if cfg.Theme.SurfaceBg != "" {
		surfaceBg = lipgloss.Color(cfg.Theme.SurfaceBg)
	}
	if cfg.Theme.Surface3 != "" {
		surface3 = lipgloss.Color(cfg.Theme.Surface3)
	}
	if cfg.Theme.Text != "" {
		textColor = lipgloss.Color(cfg.Theme.Text)
	}
	baseStyle = lipgloss.NewStyle().Background(surfaceBg).Foreground(textColor)
	statusStyle = lipgloss.NewStyle().Background(brandPrimary).Foreground(surfaceBg).Padding(0, 2).Bold(true)
	selectedStyle = lipgloss.NewStyle().Foreground(brandPrimary).Bold(true)

	// Sincronizar estilos de la barra de estado segmentada
	modeBadgeStyle = lipgloss.NewStyle().Background(brandPrimary).Foreground(surfaceBg).Padding(0, 1).Bold(true)
	infoBadgeStyle = lipgloss.NewStyle().Background(surface3).Foreground(brandPrimary).Padding(0, 1).Bold(true)
	helpBadgeStyle = lipgloss.NewStyle().Background(surfaceBg).Foreground(lipgloss.Color("#888888")).Padding(0, 1)

	// Inyectar el color de marca activo en el estilo de Glamour
	primaryHex := string(brandPrimary)
	if primaryHex == "" {
		primaryHex = "#ff007f"
	}

	if customStylePath != "" {
		customStyle, err := os.ReadFile(customStylePath)
		if err == nil {
			glamourStyle = customStyle
		} else {
			log.Printf("custom style load error %s: %v", customStylePath, err)
			glamourStyle = []byte(strings.ReplaceAll(glamourStyleTemplate, "{{brandPrimary}}", primaryHex))
		}
	} else {
		glamourStyle = []byte(strings.ReplaceAll(glamourStyleTemplate, "{{brandPrimary}}", primaryHex))
	}
}

type mode int

const (
	modeWelcome mode = iota
	modeList
	modeRead
	modeLink
)

type linkKind int

const (
	linkExternal linkKind = iota
	linkChapter
	linkAnchor
)

type link struct {
	text string
	url  string
	kind linkKind
}

type item struct {
	title    string
	filename string
	size     string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.size }
func (i item) FilterValue() string { return i.title }

type model struct {
	list            list.Model
	viewport        viewport.Model
	activeContent   string
	renderedContent string
	lang            string
	langs           []string
	langIdx         int
	archs           []string
	archIdx         int
	mode            mode
	links           []link
	linkIndex       int
	ready           bool
	width, height   int
}

func (m *model) cycleLinkNext() {
	if len(m.links) == 0 {
		return
	}
	m.linkIndex = (m.linkIndex + 1) % len(m.links)
	m.scrollToLink(m.links[m.linkIndex])
}

func (m *model) cycleLinkPrev() {
	if len(m.links) == 0 {
		return
	}
	m.linkIndex = (m.linkIndex - 1 + len(m.links)) % len(m.links)
	m.scrollToLink(m.links[m.linkIndex])
}

func (m *model) scrollToLink(l link) {
	if m.renderedContent == "" {
		return
	}
	search := l.text
	if l.kind == linkExternal && !strings.Contains(l.text, l.url) {
		search = l.url
	}
	lines := strings.Split(m.renderedContent, "\n")
	for i, line := range lines {
		if strings.Contains(line, search) {
			vh := m.viewport.Height
			if vh < 1 {
				vh = 10
			}
			half := vh / 2
			if i > half {
				m.viewport.YOffset = i - half
			} else {
				m.viewport.YOffset = 0
			}
			return
		}
	}
}

func (m *model) extractLinks(content string) {
	seen := make(map[string]bool)

	matches := markdownLinkRe.FindAllStringSubmatch(content, -1)
	m.links = make([]link, 0, len(matches)+8)

	for _, match := range matches {
		text := match[1]
		url := match[2]

		if seen[url] {
			continue
		}
		seen[url] = true

		var kind linkKind
		switch {
		case strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://"):
			kind = linkExternal
		case strings.HasPrefix(url, "#"):
			kind = linkAnchor
			url = url[1:]
		default:
			if matched, chapter := matchChapter(url); matched {
				kind = linkChapter
				url = chapter
			}
		}

		if kind == linkExternal || kind == linkChapter || kind == linkAnchor {
			m.links = append(m.links, link{text: text, url: url, kind: kind})
		}
	}

	for _, url := range rawURLRe.FindAllString(content, -1) {
		if seen[url] {
			continue
		}
		seen[url] = true
		m.links = append(m.links, link{text: url, url: url, kind: linkExternal})
	}
}

func matchChapter(url string) (bool, string) {
	parts := strings.Split(strings.TrimSuffix(url, "/"), "/")
	if len(parts) == 0 {
		return false, ""
	}
	candidate := strings.ToLower(parts[len(parts)-1])
	for _, c := range chapterOrder {
		if c == candidate {
			return true, c
		}
	}
	return false, ""
}

func clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func sidebarWidth(total int) int {
	return clamp(total*sidebarCfg.WidthPercent/100, sidebarCfg.Min, sidebarCfg.Max)
}

func openURL(url string) {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default:
		cmd = "xdg-open"
		args = []string{url}
	}
	if err := exec.Command(cmd, args...).Start(); err != nil {
		log.Printf("open url: %v", err)
	}
}

var supportedLangs = map[string]bool{
	"de": true, "nl": true, "tr": true, "es": true, "fr": true,
	"it": true, "hu": true, "pl": true, "pt": true, "sv": true,
	"cs": true, "el": true, "ru": true, "ar": true, "ta": true,
	"zh": true, "ja": true, "ko": true, "en": true,
}

func detectLang() string {
	for _, env := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		v := os.Getenv(env)
		if v == "" || v == "C" || v == "POSIX" {
			continue
		}
		parts := strings.Split(v, ".")
		lang := strings.Split(parts[0], "_")[0]
		lang = strings.ToLower(lang)
		if supportedLangs[lang] {
			return lang
		}
	}
	return "en"
}

var architectureOrder = []string{
	"amd64", "arm64", "arm", "x86", "ppc", "ppc64",
	"sparc", "mips", "alpha", "hppa",
}

var archNames = map[string]string{
	"amd64": "AMD64", "x86": "x86", "arm64": "ARM64", "arm": "ARM",
	"ppc": "PPC", "ppc64": "PPC64", "sparc": "SPARC", "mips": "MIPS",
	"alpha": "Alpha", "hppa": "HPPA",
}

var chapterOrder = []string{
	"about", "media", "networking", "disks", "stage", "base",
	"kernel", "system", "tools", "bootloader", "finalizing",
}

var chapterNames = map[string]string{
	"about":      "1.  About Gentoo",
	"media":      "2.  Installation Media",
	"networking": "3.  Configuring Network",
	"disks":      "4.  Preparing the Disks",
	"stage":      "5.  Installing Stage Files",
	"base":       "6.  Installing Base System",
	"kernel":     "7.  Configuring the Kernel",
	"system":     "8.  Configuring the System",
	"tools":      "9.  Installing System Tools",
	"bootloader": "10. Configuring the Bootloader",
	"finalizing": "11. Finalizing Installation",
}

func listAvailableArchs() []string {
	entries, err := handbookFS.ReadDir(handbookRoot)
	if err != nil {
		return []string{"amd64"}
	}
	available := make(map[string]bool, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			available[e.Name()] = true
		}
	}
	archs := make([]string, 0, len(architectureOrder))
	for _, a := range architectureOrder {
		if available[a] {
			archs = append(archs, a)
		}
	}
	if len(archs) == 0 {
		archs = []string{"amd64"}
	}
	return archs
}

func (m *model) dataDir() string {
	arch := "amd64"
	if len(m.archs) > 0 && m.archIdx < len(m.archs) {
		arch = m.archs[m.archIdx]
	}
	if m.lang == "" || m.lang == "en" {
		return getHandbookPath(arch)
	}
	return getHandbookPath(arch, m.lang)
}

func (m *model) rebuildList() {
	arch := "amd64"
	if len(m.archs) > 0 && m.archIdx < len(m.archs) {
		arch = m.archs[m.archIdx]
	}
	dataPath := m.dataDir()

	entries, err := handbookFS.ReadDir(dataPath)
	if err != nil && m.lang != "en" {
		dataPath = getHandbookPath(arch)
		entries, err = handbookFS.ReadDir(dataPath)
	}
	if err != nil {
		log.Printf("rebuild list: %v", err)
		return
	}

	entryByName := make(map[string]os.DirEntry, len(entries))
	for _, e := range entries {
		name := e.Name()
		base := strings.TrimSuffix(name, ".md")
		base = strings.TrimLeft(base, "0123456789_")
		if _, ok := entryByName[base]; !ok {
			entryByName[base] = e
		}
	}

	items := make([]list.Item, 0, len(chapterOrder))
	for _, base := range chapterOrder {
		e, ok := entryByName[base]
		if !ok {
			continue
		}
		info, err := e.Info()
		if err != nil {
			log.Printf("file info %s: %v", e.Name(), err)
			continue
		}
		title, ok := chapterNames[base]
		if !ok {
			title = strings.Title(strings.Replace(base, "_", " ", -1))
		}
		size := fmt.Sprintf("%d KB", (info.Size()+512)/1024)
		items = append(items, item{title: title, filename: e.Name(), size: size})
	}

	m.list.SetItems(items)

	archName := archNames[arch]
	if archName == "" {
		archName = strings.ToUpper(arch)
	}
	langSuffix := ""
	if m.lang != "" && m.lang != "en" {
		langSuffix = " [" + m.lang + "]"
	}
	m.list.Title = fmt.Sprintf("GENTOO HANDBOOK %s%s (%d chapters)", archName, langSuffix, len(items))
}

func (m *model) listAvailableLangs() []string {
	arch := "amd64"
	if len(m.archs) > 0 && m.archIdx < len(m.archs) {
		arch = m.archs[m.archIdx]
	}
	entries, err := handbookFS.ReadDir(getHandbookPath(arch))
	if err != nil {
		return []string{"en"}
	}
	langs := []string{"en"}
	for _, e := range entries {
		if e.IsDir() && supportedLangs[e.Name()] {
			langs = append(langs, e.Name())
		}
	}
	return langs
}

func (m *model) switchLang() {
	if len(m.langs) <= 1 {
		return
	}
	m.langIdx = (m.langIdx + 1) % len(m.langs)
	m.lang = m.langs[m.langIdx]
	m.mode = modeList
	m.activeContent = ""
	m.links = nil
	m.rebuildList()
}

func (m *model) switchArch() {
	if len(m.archs) <= 1 {
		return
	}
	m.archIdx = (m.archIdx + 1) % len(m.archs)

	prevLang := m.lang
	m.langs = m.listAvailableLangs()
	m.langIdx = 0
	for i, l := range m.langs {
		if l == prevLang {
			m.langIdx = i
			break
		}
	}
	m.lang = m.langs[m.langIdx]

	m.mode = modeList
	m.activeContent = ""
	m.links = nil
	m.rebuildList()
}

func initialModel() model {
	ensureConfig()
	cfg := loadConfig()
	applyConfig(cfg)

	archs := listAvailableArchs()

	d := list.NewDefaultDelegate()
	d.Styles.SelectedTitle = selectedStyle
	d.Styles.NormalTitle = lipgloss.NewStyle().PaddingLeft(2)

	startMode := modeWelcome
	if !showWelcome {
		startMode = modeList
	}
	m := model{archs: archs, mode: startMode, lang: "en"}
	m.langs = m.listAvailableLangs()

	if cfg.DefaultArch != "auto" {
		for i, a := range m.archs {
			if a == cfg.DefaultArch {
				m.archIdx = i
				break
			}
		}
		m.langs = m.listAvailableLangs()
	}

	if cfg.DefaultLang != "auto" {
		m.lang = cfg.DefaultLang
	} else {
		m.lang = detectLang()
	}
	m.langIdx = 0
	for i, l := range m.langs {
		if l == m.lang {
			m.langIdx = i
			break
		}
	}

	m.list = list.New(nil, d, 0, 0)
	m.rebuildList()
	m.list.SetShowStatusBar(false)
	m.list.Styles.Title = lipgloss.NewStyle().Foreground(brandPrimary).Bold(true).PaddingLeft(2)

	return m
}

func (m *model) renderChapter() {
	ww := m.viewport.Width
	if wrapWidth > 0 && wrapWidth < ww {
		ww = wrapWidth
	}
	if ww < 20 {
		ww = 80
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStylesFromJSONBytes(glamourStyle),
		glamour.WithWordWrap(ww),
		glamour.WithColorProfile(termenv.ColorProfile()),
	)
	if err != nil {
		log.Printf("glamour renderer: %v", err)
		return
	}

	str, err := renderer.Render(m.activeContent)
	if err != nil {
		log.Printf("glamour render: %v", err)
		return
	}

	m.viewport.SetContent(str)
	m.renderedContent = str
}

func (m *model) activateLink(l link) {
	switch l.kind {
	case linkExternal:
		openURL(l.url)
	case linkChapter:
		for i, base := range chapterOrder {
			if base == l.url {
				m.list.Select(i)
				m.openChapter()
				break
			}
		}
	case linkAnchor:
		lines := strings.Split(m.activeContent, "\n")
		wanted := strings.ToLower(l.url)
		for i, line := range lines {
			if strings.Contains(strings.ToLower(line), wanted) {
				if i > 3 {
					m.viewport.YOffset = i - 3
				} else {
					m.viewport.YOffset = 0
				}
				break
			}
		}
	}
}

func (m *model) openChapter() {
	i, ok := m.list.SelectedItem().(item)
	if !ok {
		return
	}

	arch := "amd64"
	if len(m.archs) > 0 && m.archIdx < len(m.archs) {
		arch = m.archs[m.archIdx]
	}
	data, err := handbookFS.ReadFile(strings.Join([]string{m.dataDir(), i.filename}, "/"))
	if err != nil && m.lang != "en" {
		data, err = handbookFS.ReadFile(getHandbookPath(arch, i.filename))
	}
	if err != nil {
		log.Printf("read chapter %s: %v", i.filename, err)
		return
	}

	m.activeContent = langLinkRe.ReplaceAllString(string(data), "")
	m.activeContent = editLinkRe.ReplaceAllString(m.activeContent, " ")
	m.extractLinks(m.activeContent)
	m.renderChapter()
	m.viewport.GotoTop()
	m.mode = modeRead
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.mode == modeWelcome {
			if msg.String() == "q" || msg.String() == "ctrl+c" {
				return m, tea.Quit
			}
			m.list.Select(0)
			m.openChapter()
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "tab":
			if m.mode == modeRead && len(m.links) > 0 {
				m.mode = modeLink
				m.scrollToLink(m.links[m.linkIndex])
				return m, nil
			}
			if m.mode == modeLink {
				m.cycleLinkNext()
			}

		case "shift+tab":
			if m.mode == modeLink {
				m.cycleLinkPrev()
			}

		case "up", "k":
			if m.mode == modeLink {
				m.cycleLinkPrev()
			} else if m.mode == modeRead {
				m.viewport.LineUp(1)
			}

		case "down", "j":
			if m.mode == modeLink {
				m.cycleLinkNext()
			} else if m.mode == modeRead {
				m.viewport.LineDown(1)
			}

		case "a":
			if m.mode == modeList || m.mode == modeRead {
				m.switchArch()
			}

		case "L":
			if m.mode == modeList || m.mode == modeRead {
				m.switchLang()
			}

		case "l", "right":
			if m.mode == modeList {
				m.openChapter()
			} else if m.mode == modeRead && len(m.links) > 0 {
				m.mode = modeLink
				m.scrollToLink(m.links[m.linkIndex])
			}

		case "enter":
			if m.mode == modeList {
				m.openChapter()
			} else if m.mode == modeLink && len(m.links) > 0 {
				m.activateLink(m.links[m.linkIndex])
			}

		case "o":
			if m.mode == modeLink && len(m.links) > 0 && m.links[m.linkIndex].kind == linkExternal {
				openURL(m.links[m.linkIndex].url)
			}

		case "left", "h", "esc":
			if m.mode == modeLink {
				m.mode = modeRead
			} else if m.mode == modeRead {
				m.mode = modeList
				m.activeContent = ""
				m.links = nil
			}
		}

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		lw := sidebarWidth(m.width)
		vh := m.height - 2
		vpWidth := m.width - lw - 1
		if vpWidth < 1 {
			vpWidth = 1
		}
		m.list.SetSize(lw, vh)
		m.viewport = viewport.New(vpWidth, vh)
		if m.activeContent != "" {
			m.renderChapter()
		}
		m.ready = true
	}

	if m.mode == modeList {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		if cmd != nil {
			return m, cmd
		}
	}
	if m.mode == modeRead && m.activeContent != "" {
		var cmd tea.Cmd
		m.viewport, cmd = m.viewport.Update(msg)
		if cmd != nil {
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {
	if !m.ready {
		return ""
	}
	if m.mode == modeWelcome {
		return m.welcomeView()
	}

	loc := getUILocale(m.lang)

	lw := sidebarWidth(m.width)
	listRendered := m.list.View()

	borderRendered := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(surface3).
		Render("")

	vpWidth := m.width - lw - 1
	if vpWidth < 1 {
		vpWidth = 1
	}
	viewStyle := lipgloss.NewStyle().Width(vpWidth)

	var contentStr string
	if m.activeContent == "" {
		contentStr = lipgloss.NewStyle().
			Width(vpWidth).
			Foreground(lipgloss.Color("#666666")).
			Render(loc.selectChapter)
	} else {
		contentStr = viewStyle.Render(m.viewport.View())
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, listRendered, borderRendered, contentStr)

	var status string
	statusWidth := m.width
	if statusWidth < 10 {
		statusWidth = 80
	}

	var modePart, infoPart, helpPart string

	if m.mode == modeList {
		modePart = modeBadgeStyle.Render(" LIST ")

		arch := "amd64"
		if len(m.archs) > 0 && m.archIdx < len(m.archs) {
			arch = m.archs[m.archIdx]
		}
		archName := archNames[arch]
		if archName == "" {
			archName = strings.ToUpper(arch)
		}
		lbl := m.lang
		if lbl == "en" {
			lbl = "EN"
		}
		infoPart = infoBadgeStyle.Render(fmt.Sprintf(" %s │ %s ", strings.ToUpper(archName), strings.ToUpper(lbl)))

		helpPart = helpBadgeStyle.Render(fmt.Sprintf(" ↑↓ %s  \u2022  Enter/l %s  \u2022  q %s ", loc.helpNavigate, loc.helpOpen, loc.helpQuit))
	} else if m.mode == modeRead {
		modePart = modeBadgeStyle.Render(" READ ")

		var title string
		if item, ok := m.list.SelectedItem().(item); ok {
			title = item.title
			if len(title) > 25 {
				title = title[:22] + "..."
			}
		} else {
			title = "Handbook"
		}
		infoPart = infoBadgeStyle.Render(fmt.Sprintf(" %s ", title))

		linkHelp := ""
		if len(m.links) > 0 {
			linkHelp = fmt.Sprintf("  \u2022  l/Tab %s", loc.helpLinks)
		}
		helpPart = helpBadgeStyle.Render(fmt.Sprintf(" ↑↓/jk %s  \u2022  h/← %s%s  \u2022  q %s ", loc.helpScroll, loc.helpBack, linkHelp, loc.helpQuit))
	} else if len(m.links) > 0 {
		modePart = modeBadgeStyle.Render(" LINKS ")

		l := m.links[m.linkIndex]
		kindLabel := "URL"
		if l.kind == linkChapter {
			kindLabel = "Chapter"
		} else if l.kind == linkAnchor {
			kindLabel = "Anchor"
		}
		label := l.text
		if len(label) > 30 {
			label = label[:27] + "..."
		}
		infoPart = infoBadgeStyle.Render(fmt.Sprintf(" %d/%d │ [%s] %s ", m.linkIndex+1, len(m.links), kindLabel, label))

		helpPart = helpBadgeStyle.Render(fmt.Sprintf(" ↑↓/Tab %s  \u2022  Enter %s  \u2022  h/Esc %s  \u2022  q %s ", loc.helpCycle, loc.helpOpen, loc.helpBack, loc.helpQuit))
	}

	leftBar := lipgloss.JoinHorizontal(lipgloss.Top, modePart, infoPart)

	leftWidth := lipgloss.Width(leftBar)
	helpWidth := lipgloss.Width(helpPart)
	space := statusWidth - leftWidth - helpWidth
	if space < 0 {
		space = 0
	}

	spacer := lipgloss.NewStyle().Background(surfaceBg).Render(strings.Repeat(" ", space))

	status = lipgloss.JoinHorizontal(lipgloss.Top, leftBar, spacer, helpPart)
	status = lipgloss.NewStyle().Background(surfaceBg).Width(statusWidth).Render(status)

	return baseStyle.Render(lipgloss.JoinVertical(lipgloss.Top, row, status))
}

type uiLocales struct {
	welcomeTitle      string
	welcomeLastFetch  string
	welcomeShortcuts  string
	welcomeNav        string
	welcomeOpen       string
	welcomeBack       string
	welcomeLinks      string
	welcomeArch       string
	welcomeLang       string
	welcomeQuit       string
	welcomePrompt     string

	selectChapter     string
	helpNavigate      string
	helpOpen          string
	helpQuit          string
	helpScroll        string
	helpBack          string
	helpLinks         string
	helpCycle         string
}

var uiTranslations = map[string]uiLocales{
	"en": {
		welcomeTitle:      "Gentoo Linux Installation Manual in your Terminal",
		welcomeLastFetch:  "Last Fetch",
		welcomeShortcuts:  "QUICK KEYBOARD SHORTCUTS",
		welcomeNav:        " j / k / ↑ / ↓  : Navigate chapters and scroll",
		welcomeOpen:       " Enter / l / →  : Open chapter or activate link",
		welcomeBack:       " h / ← / Esc    : Go back (Reading / List)",
		welcomeLinks:      " Tab / Shift+Tab: Cycle links in manual",
		welcomeArch:       " a              : Change architecture (amd64, arm64, etc.)",
		welcomeLang:       " L (Shift+l)    : Change chapter language",
		welcomeQuit:       " q / Ctrl+C     : Exit application",
		welcomePrompt:     "Press any key to start reading...",

		selectChapter:     "\n\n  Select a chapter and press Enter\n\n  \u2191\u2193 navigate  \u00b7  Enter open  \u00b7  q quit",
		helpNavigate:      "navigate",
		helpOpen:          "open",
		helpQuit:          "quit",
		helpScroll:        "scroll",
		helpBack:          "back",
		helpLinks:         "links",
		helpCycle:         "cycle",
	},
	"es": {
		welcomeTitle:      "Manual de Instalación de Gentoo Linux en tu Terminal",
		welcomeLastFetch:  "Último Fetch",
		welcomeShortcuts:  "ATAJOS DE TECLADO RÁPIDOS",
		welcomeNav:        " j / k / ↑ / ↓  : Navegar capítulos y scroll",
		welcomeOpen:       " Enter / l / →  : Abrir capítulo o activar enlace",
		welcomeBack:       " h / ← / Esc    : Volver atrás (Lectura / Lista)",
		welcomeLinks:      " Tab / Shift+Tab: Ciclar enlaces en manual",
		welcomeArch:       " a              : Cambiar arquitectura (amd64, arm64, etc.)",
		welcomeLang:       " L (Shift+l)    : Cambiar idioma de los capítulos",
		welcomeQuit:       " q / Ctrl+C     : Salir de la aplicación",
		welcomePrompt:     "Presiona cualquier tecla para comenzar a leer...",

		selectChapter:     "\n\n  Selecciona un capítulo y presiona Enter\n\n  \u2191\u2193 navegar  \u00b7  Enter abrir  \u00b7  q salir",
		helpNavigate:      "navegar",
		helpOpen:          "abrir",
		helpQuit:          "salir",
		helpScroll:        "desplazar",
		helpBack:          "atrás",
		helpLinks:         "enlaces",
		helpCycle:         "ciclar",
	},
}

func getUILocale(lang string) uiLocales {
	if loc, ok := uiTranslations[lang]; ok {
		return loc
	}
	return uiTranslations["en"]
}

func (m model) welcomeView() string {
	logo := `______     ______     __   __     ______   ______     ______          
/\  ___\   /\  ___\   /\ "-.\ \   /\__  _\ /\  __ \   /\  __ \         
\ \ \__ \  \ \  __\   \ \ \-.  \  \/_/\ \/ \ \ \/\ \  \ \ \/\ \        
 \ \_____\  \ \_____\  \ \_\\"\_\    \ \_\  \ \_____\  \ \_____\       
  \/_____/   \/_____/   \/_/ \/_/     \/_/   \/_____/   \/_____/       
                                                                       
 ______   __  __     __     ______     ______     ______     __  __    
/\__  _\ /\ \/\ \   /\ \   /\  == \   /\  __ \   /\  __ \   /\ \/ /    
\/_/\ \/ \ \ \_\ \  \ \ \  \ \  __<   \ \ \/\ \  \ \ \/\ \  \ \  _"-.  
   \ \_\  \ \_____\  \ \_\  \ \_____\  \ \_____\  \ \_____\  \ \_\ \_\ 
    \/_/   \/_____/   \/_/   \/_____/   \/_____/   \/_____/   \/_/\/_/`

	logoStyled := lipgloss.NewStyle().Foreground(brandPrimary).Bold(true).Render(logo)

	loc := getUILocale(m.lang)

	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true).
		Render(loc.welcomeTitle)

	versionStr := "Version: v0.1.0"
	fetchStr := fmt.Sprintf("%s: %s", loc.welcomeLastFetch, getLastFetchDate())
	info := lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Render(fmt.Sprintf("%s  \u2022  %s", versionStr, fetchStr))

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(surface3).
		Padding(1, 4).
		Width(65).
		Align(lipgloss.Center)

	shortcuts := fmt.Sprintf("%s\n\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		loc.welcomeShortcuts,
		loc.welcomeNav,
		loc.welcomeOpen,
		loc.welcomeBack,
		loc.welcomeLinks,
		loc.welcomeArch,
		loc.welcomeLang,
		loc.welcomeQuit,
	)

	shortcutsStyled := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A3A3A3")).
		Render(shortcuts)

	box := boxStyle.Render(shortcutsStyled)

	prompt := lipgloss.NewStyle().
		Foreground(brandPrimary).
		Bold(true).
		Render(loc.welcomePrompt)

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		"",
		logoStyled,
		"",
		title,
		info,
		"",
		box,
		"",
		"",
		prompt,
		"",
	)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content, lipgloss.WithWhitespaceBackground(surfaceBg))
}

func getLastFetchDate() string {
	data, err := handbookFS.ReadFile(getHandbookPath("last_fetch.txt"))
	if err != nil {
		return defaultFetchDate
	}
	return strings.TrimSpace(string(data))
}

func initLogging() *os.File {
	ensureConfig()
	logPath := filepath.Join(configDir(), "debug.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil
	}
	log.SetOutput(file)
	log.Println("--- gentoo-tuibook session started ---")
	return file
}

func main() {
	logFile := initLogging()
	if logFile != nil {
		defer logFile.Close()
	}
	tea.NewProgram(initialModel(), tea.WithAltScreen()).Run()
}

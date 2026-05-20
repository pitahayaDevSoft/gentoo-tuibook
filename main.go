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
	editLinkRe  = regexp.MustCompile(`\\\[[ \t]*\[[ \t]*[Ee]dit[ \t]*\][ \t]*\([^)]*\)[ \t]*\\\]`)
	langLinkRe  = regexp.MustCompile(`(?m)^Other languages:\n\n(?:- .*\n)+\n`)

	glamourStyle = []byte(`{
		"document": { "block_prefix": "\n", "block_suffix": "\n", "color": "#EBEBEB", "margin": 1 },
		"heading": { "block_suffix": "\n", "color": "#22d3ee", "bold": true },
		"h1": { "prefix": " ", "suffix": " ", "color": "#22d3ee", "background_color": "#1c2130", "bold": true },
		"h2": { "prefix": "\u2554 ", "color": "#22d3ee", "bold": true },
		"h3": { "prefix": "\u2551 ", "color": "#22d3ee" },
		"strong": { "bold": true, "color": "#FFFFFF" },
		"emph": { "italic": true, "color": "#94A3B8" },
		"link": { "color": "#22d3ee", "underline": true },
		"link_text": { "color": "#22d3ee", "bold": true },
		"code": { "prefix": " ", "suffix": " ", "color": "#F0C674", "background_color": "#1C2130" },
		"code_block": { "margin": 1, "chroma": {
			"text": { "color": "#C4C4C4" },
			"keyword": { "color": "#22D3EE" },
			"name_function": { "color": "#7FDBCA" },
			"literal_string": { "color": "#F0C674" },
			"comment": { "color": "#6272A4" },
			"name_builtin": { "color": "#7FDBCA" },
			"operator": { "color": "#EF8080" },
			"literal_number": { "color": "#6EEFC0" },
			"keyword_type": { "color": "#B294BB" },
			"name_class": { "color": "#FFFFFF", "underline": true, "bold": true },
			"name_attribute": { "color": "#7A7AE6" },
			"name_tag": { "color": "#B083EA" },
			"background": { "background_color": "#121620" }
		}},
		"item": { "block_prefix": "\u2022 ", "color": "#EBEBEB" },
		"enumeration": { "block_prefix": ". " },
		"hr": { "color": "#22D3EE", "format": "\n\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\u2501\n" },
		"block_quote": { "color": "#6272A4", "indent": 1, "indent_token": "\u2502 " },
		"table": { "center_separator": "\u253C", "column_separator": "\u2502", "row_separator": "\u2500" }
	}`)

	surfaceBg    = lipgloss.Color("#060608")
	surface3     = lipgloss.Color("#1c2130")
	brandPrimary = lipgloss.Color("#22d3ee")

	baseStyle     = lipgloss.NewStyle().Background(surfaceBg).Foreground(lipgloss.Color("#EBEBEB"))
	statusStyle   = lipgloss.NewStyle().Background(brandPrimary).Foreground(surfaceBg).Padding(0, 2).Bold(true)
	selectedStyle = lipgloss.NewStyle().Foreground(brandPrimary).Bold(true)

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
	DefaultArch string        `json:"default_arch"`
	DefaultLang string        `json:"default_lang"`
	Sidebar     SidebarConfig `json:"sidebar"`
	Theme       ThemeColors   `json:"theme"`
}

func defaultConfig() Config {
	return Config{
		DefaultArch: "auto",
		DefaultLang: "auto",
		Sidebar:     SidebarConfig{WidthPercent: 25, Min: 20, Max: 40},
		Theme: ThemeColors{
			BrandPrimary: "#22d3ee",
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
}

type mode int

const (
	modeList mode = iota
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

	re := regexp.MustCompile(`\[([^\]]+)\]\(([^\s\)]+)(?:\s+"[^"]*")?\)`)
	matches := re.FindAllStringSubmatch(content, -1)
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

	re2 := regexp.MustCompile(`https?://[^\s\)"]+`)
	for _, url := range re2.FindAllString(content, -1) {
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
	entries, err := handbookFS.ReadDir("data/handbook")
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
		return "data/handbook/" + arch
	}
	return "data/handbook/" + arch + "/" + m.lang
}

func (m *model) rebuildList() {
	arch := "amd64"
	if len(m.archs) > 0 && m.archIdx < len(m.archs) {
		arch = m.archs[m.archIdx]
	}
	dataPath := "data/handbook/" + arch
	if m.lang != "" && m.lang != "en" {
		dataPath += "/" + m.lang
	}

	entries, err := handbookFS.ReadDir(dataPath)
	if err != nil && m.lang != "en" {
		dataPath = "data/handbook/" + arch
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
	entries, err := handbookFS.ReadDir("data/handbook/" + arch)
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

	m := model{archs: archs, mode: modeList, lang: "en"}
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
	if ww < 20 {
		ww = 80
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStylesFromJSONBytes(glamourStyle),
		glamour.WithWordWrap(ww),
		glamour.WithColorProfile(termenv.TrueColor),
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
	data, err := handbookFS.ReadFile(m.dataDir() + "/" + i.filename)
	if err != nil && m.lang != "en" {
		data, err = handbookFS.ReadFile("data/handbook/" + arch + "/" + i.filename)
	}
	if err != nil {
		log.Printf("read chapter %s: %v", i.filename, err)
		return
	}

	m.activeContent = langLinkRe.ReplaceAllString(string(data), "")
	m.activeContent = editLinkRe.ReplaceAllString(m.activeContent, "")
	m.extractLinks(m.activeContent)
	m.renderChapter()
	m.viewport.GotoTop()
	m.mode = modeRead
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
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
			}

		case "down", "j":
			if m.mode == modeLink {
				m.cycleLinkNext()
			}

		case "a":
			if m.mode == modeList || m.mode == modeRead {
				m.switchArch()
			}

		case "L":
			if m.mode == modeList || m.mode == modeRead {
				m.switchLang()
			}

		case "l":
			if m.mode == modeRead && len(m.links) > 0 {
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

		case "left", "esc":
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

	archHint := ""
	if len(m.archs) > 1 {
		archHint = " \u2502 a arch"
	}
	langHint := ""
	if len(m.langs) > 1 {
		lbl := m.lang
		if lbl == "en" {
			lbl = "EN"
		}
		langHint = fmt.Sprintf(" \u2502 L %s", lbl)
	}

	var contentStr string
	if m.activeContent == "" {
		contentStr = lipgloss.NewStyle().
			Width(vpWidth).
			Foreground(lipgloss.Color("#666666")).
			Render("\n\n  Select a chapter and press Enter\n\n  \u2191\u2193 navigate  \u00b7  Enter open  \u00b7  q quit")
	} else {
		contentStr = viewStyle.Render(m.viewport.View())
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, listRendered, borderRendered, contentStr)

	var status string
	if m.mode == modeList {
		status = statusStyle.Render(fmt.Sprintf(" [ LIST ] \u2191\u2193 navigate%s%s \u2502  Enter open  \u2502  q quit ", langHint, archHint))
	} else if m.mode == modeRead {
		linkHint := ""
		if len(m.links) > 0 {
			linkHint = " \u2502 l links"
		}
		status = statusStyle.Render(fmt.Sprintf(" [ READ ] \u2191\u2193 scroll%s%s%s \u2502 \u2190 back  \u2502 q quit ", linkHint, langHint, archHint))
	} else if len(m.links) > 0 {
		l := m.links[m.linkIndex]
		kindLabel := "URL"
		if l.kind == linkChapter {
			kindLabel = "Chapter"
		} else if l.kind == linkAnchor {
			kindLabel = "Anchor"
		}
		label := l.text
		if len(label) > 40 {
			label = label[:37] + "..."
		}
		status = statusStyle.Render(fmt.Sprintf(" [ LINKS %d/%d ] [%s] %s \u2502 \u2191\u2193/Tab  \u2502 Enter open  \u2502 Esc back ", m.linkIndex+1, len(m.links), kindLabel, label))
	}

	return baseStyle.Render(lipgloss.JoinVertical(lipgloss.Top, row, status))
}

func main() {
	tea.NewProgram(initialModel(), tea.WithAltScreen()).Run()
}

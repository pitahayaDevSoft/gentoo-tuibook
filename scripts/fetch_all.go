package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type WikiResponse struct {
	Parse struct {
		Text struct {
			Content string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
	Error *struct {
		Code string `json:"code"`
	} `json:"error"`
}

var architectures = []string{
	"amd64", "arm64", "arm", "x86", "ppc", "ppc64",
	"sparc", "mips", "alpha", "hppa",
}

var wikiArch = map[string]string{
	"amd64": "AMD64", "x86": "X86", "arm64": "ARM64", "arm": "ARM",
	"ppc": "PPC", "ppc64": "PPC64", "sparc": "SPARC", "mips": "MIPS",
	"alpha": "Alpha", "hppa": "HPPA",
}

var chapters = []string{
	"About", "Media", "Networking", "Disks", "Stage", "Base",
	"Kernel", "System", "Tools", "Bootloader", "Finalizing",
}

var languages = []string{
	"de", "nl", "tr", "es", "fr", "it", "hu", "pl", "pt", "sv",
	"cs", "el", "ru", "ar", "ta", "zh", "ja", "ko",
}

var rateLimit = 500 * time.Millisecond

func pageName(arch, chapter, lang string) string {
	page := fmt.Sprintf("Handbook:%s/Installation/%s", wikiArch[arch], chapter)
	if lang != "en" {
		page += "/" + lang
	}
	return page
}

func fetchPage(client *http.Client, page string) (*WikiResponse, error) {
	url := "https://wiki.gentoo.org/api.php?action=parse&page=" + page + "&format=json"
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	var apiResp WikiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("json: %w", err)
	}
	return &apiResp, nil
}

func main() {
	converter := md.NewConverter("", true, nil)
	client := &http.Client{Timeout: 30 * time.Second}

	total, success, skipped, fails := 0, 0, 0, 0

	for _, arch := range architectures {
		langs := append([]string{""}, languages...)
		for _, lang := range langs {
			langCode := lang
			if langCode == "" {
				langCode = "en"
			}

			outDir := filepath.Join("data", "handbook", arch)
			if lang != "" {
				outDir = filepath.Join(outDir, lang)
			}
			if err := os.MkdirAll(outDir, 0755); err != nil {
				fmt.Printf("ERROR mkdir %s: %v\n", outDir, err)
				continue
			}

			for _, chapter := range chapters {
				filename := strings.ToLower(chapter) + ".md"
				path := filepath.Join(outDir, filename)

				if _, err := os.Stat(path); err == nil {
					skipped++
					continue
				}

				page := pageName(arch, chapter, langCode)
				total++

				fmt.Printf("[%02d/%02d] %-45s ", total, len(architectures)*len(langs)*len(chapters), page)

				apiResp, err := fetchPage(client, page)
				if err != nil {
					fmt.Printf("FAIL %v\n", err)
					fails++
					time.Sleep(2 * time.Second)
					continue
				}

				if apiResp.Error != nil || apiResp.Parse.Text.Content == "" {
					fmt.Printf("MISS\n")
					time.Sleep(rateLimit)
					continue
				}

				markdown, err := converter.ConvertString(apiResp.Parse.Text.Content)
				if err != nil {
					fmt.Printf("FAIL convert: %v\n", err)
					fails++
					time.Sleep(2 * time.Second)
					continue
				}

				content := "# " + chapter + "\n\n" + markdown
				if err := os.WriteFile(path, []byte(content), 0644); err != nil {
					fmt.Printf("FAIL write: %v\n", err)
					fails++
					time.Sleep(2 * time.Second)
					continue
				}

				success++
				fmt.Printf("OK\n")
				time.Sleep(rateLimit)
			}
		}
	}

	fmt.Printf("\n=== Done: %d downloaded, %d skipped, %d failed ===\n", success, skipped, fails)
}

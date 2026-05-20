package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type WikiResponse struct {
	Parse struct {
		Text struct {
			Content string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
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

func pageName(arch, chapter string) string {
	return fmt.Sprintf("Handbook:%s/Installation/%s", wikiArch[arch], chapter)
}

func main() {
	converter := md.NewConverter("", true, nil)

	for _, arch := range architectures {
		outDir := filepath.Join("..", "data", "handbook", arch)
		if err := os.MkdirAll(outDir, 0755); err != nil {
			fmt.Printf("ERROR creating dir %s: %v\n", outDir, err)
			continue
		}

		for _, chapter := range chapters {
			page := pageName(arch, chapter)
			filename := strings.ToLower(chapter) + ".md"
			path := filepath.Join(outDir, filename)

			fmt.Printf("Fetching %s...\n", page)

			resp, err := http.Get("https://wiki.gentoo.org/api.php?action=parse&page=" + page + "&format=json")
			if err != nil {
				fmt.Printf("  HTTP error: %v\n", err)
				continue
			}
			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Printf("  read error: %v\n", err)
				continue
			}

			var apiResp WikiResponse
			if err := json.Unmarshal(body, &apiResp); err != nil {
				fmt.Printf("  JSON error: %v\n", err)
				continue
			}

			markdown, err := converter.ConvertString(apiResp.Parse.Text.Content)
			if err != nil {
				fmt.Printf("  convert error: %v\n", err)
				continue
			}

			content := "# " + chapter + "\n\n" + markdown
			if err := os.WriteFile(path, []byte(content), 0644); err != nil {
				fmt.Printf("  write error: %v\n", err)
				continue
			}

			fmt.Printf("  -> %s\n", path)
		}
	}
}

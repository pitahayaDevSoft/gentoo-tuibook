package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

type WikiResponse struct {
	Parse struct {
		Title string `json:"title"`
		Text  struct {
			Content string `json:"*"`
		} `json:"text"`
	} `json:"parse"`
}

var pages = map[string]string{
	"01_about":   "Handbook:AMD64/Installation/About",
	"02_media":   "Handbook:AMD64/Installation/Media",
	"03_network": "Handbook:AMD64/Installation/Networking",
	"04_disks":   "Handbook:AMD64/Installation/Disks",
}

func main() {
	converter := md.NewConverter("", true, nil)
	outDir := filepath.Join("..", "data", "handbook", "amd64")

	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("Failed to create dir: %v", err)
	}

	for filename, title := range pages {
		fmt.Printf("Fetching %s...\n", title)
		url := fmt.Sprintf("https://wiki.gentoo.org/api.php?action=parse&page=%s&format=json", title)
		
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "GentooTuiBook/1.0")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("Error fetching %s: %v", title, err)
			continue
		}
		
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("Error reading body: %v", err)
			continue
		}

		var apiResp WikiResponse
		if err := json.Unmarshal(body, &apiResp); err != nil {
			log.Printf("Error parsing JSON: %v", err)
			continue
		}

		htmlContent := apiResp.Parse.Text.Content
		if htmlContent == "" {
			continue
		}

		markdown, err := converter.ConvertString(htmlContent)
		if err != nil {
			log.Printf("Error converting HTML: %v", err)
			continue
		}

		finalMarkdown := fmt.Sprintf("# %s\n\n%s", title, markdown)
		outPath := filepath.Join(outDir, filename+".md")
		if err := os.WriteFile(outPath, []byte(finalMarkdown), 0644); err != nil {
			log.Printf("Error writing file: %v", err)
		} else {
			fmt.Printf("Saved %s\n", outPath)
		}
	}
}
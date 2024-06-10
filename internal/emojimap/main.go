package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	URL              = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"
	OUTPUT_FILE_NAME = "emoji_map.json"
)

type Emoji struct {
	Emoji       string   `json:"emoji"`
	Description string   `json:"description"`
	Aliases     []string `json:"aliases"`
	Tags        []string `json:"tags"`
}

func main() {
	outputPath := flag.String("output-path", "", "defines where the emoji map will be stored")
	flag.Parse()
	generateMap(*outputPath)
}

func generateMap(outputPath string) {
	if !isOutputPathValid(outputPath) {
		log.Fatal("'-output-path' flag is required")
	}
	emojiMap := generateEmojiMap(URL)
	storeMapToJson(emojiMap, outputPath)
}

func isOutputPathValid(outputPath string) bool {
	return outputPath != ""
}

func storeMapToJson(emojiMap map[string][]string, filePath string) {
	data, err := json.MarshalIndent(emojiMap, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("emoji map generated and stored at: %s\n", filePath)
}

func generateEmojiMap(url string) map[string][]string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error fetching data: %v\n", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v\n", err)
	}

	var emojis []Emoji
	err = json.Unmarshal(data, &emojis)
	if err != nil {
		log.Fatalf("error unmarshaling JSON: %v\n", err)
	}

	emojiMap := make(map[string][]string)
	for _, emoji := range emojis {
		addToMap(emojiMap, strings.ToLower(emoji.Description), emoji.Emoji, true)
		for _, alias := range emoji.Aliases {
			addToMap(emojiMap, strings.ToLower(alias), emoji.Emoji, true)
		}
		for _, tag := range emoji.Tags {
			addToMap(emojiMap, strings.ToLower(tag), emoji.Emoji, false)
		}
	}

	return emojiMap
}

func addToMap(m map[string][]string, key, emoji string, prepend bool) {
	sanitizedKey := strings.ReplaceAll(key, "_", " ")

	if _, ok := m[sanitizedKey]; !ok {
		m[sanitizedKey] = []string{}
	}

	for _, alias := range m[sanitizedKey] {
		if alias == emoji {
			return
		}
	}

	if prepend {
		m[sanitizedKey] = append([]string{emoji}, m[sanitizedKey]...)
	} else {
		m[sanitizedKey] = append(m[sanitizedKey], emoji)
	}
}

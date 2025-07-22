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
	defaultURL      = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"
	outputFileName  = "emoji_map.json"
	filePermissions = 0600
	maxResponseSize = 10 << 20 // 10MB
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
	emojiMap := generateEmojiMap(defaultURL)
	storeMapToJSON(emojiMap, outputPath)
	log.Printf("emoji map generated and stored at: %s\n", outputPath)

	maxKeyLength, longestKey := getMaxWordsInKey(emojiMap)
	log.Printf("longest key '%s' was '%d' words long\n", longestKey, maxKeyLength)
}

func isOutputPathValid(outputPath string) bool {
	return outputPath != ""
}

func storeMapToJSON(emojiMap map[string][]string, filePath string) {
	data, err := json.MarshalIndent(emojiMap, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filePath, data, filePermissions)
	if err != nil {
		log.Fatal(err)
	}
}

func generateEmojiMap(url string) map[string][]string {
	resp, err := http.Get(url) //nolint:noctx // This is a CLI tool, context not needed
	if err != nil {
		log.Fatalf("error fetching data: %v\n", err)
	}

	// Check HTTP status code for security
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("HTTP request failed with status: %d", resp.StatusCode)
	}

	// Limit response size to prevent memory exhaustion attacks
	limitedReader := io.LimitReader(resp.Body, maxResponseSize)

	data, err := io.ReadAll(limitedReader)
	resp.Body.Close() // Close immediately after reading
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

func getMaxWordsInKey(emojiMap map[string][]string) (maxWords int, longestKey string) {
	maxWords = 0
	longestKey = ""

	for key := range emojiMap {
		words := strings.Split(key, " ")
		numberOfWords := len(words)
		if numberOfWords > maxWords {
			maxWords = numberOfWords
			longestKey = key
		}
	}
	return maxWords, longestKey
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

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	URL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"
)

type Emoji struct {
	Emoji       string   `json:"emoji"`
	Description string   `json:"description"`
	Aliases     []string `json:"aliases"`
	Tags        []string `json:"tags"`
}

func main() {
	emojiMap := generateEmojiMap()
	storeMapToJson(emojiMap, getFilePathOfExecutable())
}

func storeMapToJson(emojiMap map[string][]string, filePath string) {
	filePath = filepath.Join(filePath, "emoji_map.json")
	data, err := json.MarshalIndent(emojiMap, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Emoji map generated and stored at: %s\n", filePath)
}

func getFilePathOfExecutable() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func generateEmojiMap() map[string][]string {
	resp, err := http.Get("https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json")
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil
	}

	var emojis []Emoji
	err = json.Unmarshal(data, &emojis)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return nil
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

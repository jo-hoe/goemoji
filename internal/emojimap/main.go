package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	URL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"
)

type emoji struct {
	Emoji     string   `json:"emoji"`
	Description string `json:"description"`
	Aliases   []string `json:"aliases"`
	Tags      []string `json:"tags"`
}


// creates the json map based on
// https://github.com/github/gemoji/blob/master/db/emoji.json
func main() {
	// create emoji map
	emojis := make([]emoji, 0)

	response, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}

	// unmashal json response and put it in a map with. maps keys is a word(s) from the field
	// "description", "aliases", or "tags" values is a list of emoji(s)
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&emojis)
	if err != nil {
		log.Fatalln(err)
	}

	// create a map of emoji to their description, aliases, and tags
	emojiMap := make(map[string][]string)
	for _, e := range emojis {
		// add emojies of descriptions and aliases in front
		// convert aliases to remove "_" to " "
		// add tags last
	}

	// store map in json

}

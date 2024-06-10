package goemoji

import (
	"embed"
	"encoding/json"
)

//go:embed emoji_map.json
var emojiFileSystem embed.FS

const (
	embedFileName = "emoji_map.json"
)

type Emojifier struct {
	strategy EmojifyStrategy
	emojiMap map[string][]string
}

func NewEmojifier(strategy EmojifyStrategy) (*Emojifier, error) {
	loadedMap, err := loadEmojiMap()
	if err != nil {
		return nil, err
	}

	return &Emojifier{
		strategy: strategy,
		emojiMap: loadedMap,
	}, nil
}

func (e *Emojifier) Emojify(text string) string {
	return e.strategy.Emojify(text, e.emojiMap)
}

func loadEmojiMap() (emojiMap map[string][]string, err error) {
	data, err := emojiFileSystem.ReadFile(embedFileName)
	if err != nil {
		return emojiMap, err
	}
	if err := json.Unmarshal(data, &emojiMap); err != nil {
		return emojiMap, err
	}

	return emojiMap, nil
}

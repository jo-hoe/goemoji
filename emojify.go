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
	strategy          EmojifyStrategy
	emojiTags         map[string][]string
	emojiSet          map[string]bool
	minimumWordLength int
}

func NewDefaultEmojifier() (*Emojifier, error) {
	return NewEmojifier(ReplaceSubstring{}, 4)
}

func NewEmojifier(strategy EmojifyStrategy, minimumWordLength int) (*Emojifier, error) {
	loadedMap, err := loadEmojiMap()
	if err != nil {
		return nil, err
	}

	return &Emojifier{
		strategy:          strategy,
		emojiTags:         loadedMap,
		emojiSet:          createEmojiSet(loadedMap),
		minimumWordLength: minimumWordLength,
	}, nil
}

func (e *Emojifier) Emojify(text string) string {
	return e.strategy.Emojify(text, e.minimumWordLength, e.emojiTags, e.emojiSet)
}

func (e *Emojifier) ContainsEmoji(text string) bool {
	emojis := e.ExtractEmojis(text)
	return len(emojis) > 0
}

func (e *Emojifier) ExtractEmojis(text string) []string {
	return extractEmojis(text, e.emojiSet)
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

func createEmojiSet(emojiTags map[string][]string) map[string]bool {
	result := make(map[string]bool, 0)
	for _, emojis := range emojiTags {
		for _, emoji := range emojis {
			result[emoji] = true
		}
	}
	return result
}

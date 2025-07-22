package goemoji

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed emoji_map.json
var emojiFileSystem embed.FS

const (
	embedFileName        = "emoji_map.json"
	defaultMinWordLength = 4
)

// Emojifier provides functionality to add emojis to text using different strategies.
// It is safe for concurrent use by multiple goroutines.
type Emojifier struct {
	strategy          EmojifyStrategy
	emojiTags         map[string][]string
	emojiSet          map[string]bool
	minimumWordLength int
}

// NewDefaultEmojifier creates a new Emojifier with default settings.
// It uses ReplaceSubstring strategy and minimum word length of 4.
func NewDefaultEmojifier() (*Emojifier, error) {
	return NewEmojifier(ReplaceSubstring{}, defaultMinWordLength)
}

// NewEmojifier creates a new Emojifier with the specified strategy and minimum word length.
// Returns an error if strategy is nil or minimumWordLength is negative.
func NewEmojifier(strategy EmojifyStrategy, minimumWordLength int) (*Emojifier, error) {
	if strategy == nil {
		return nil, fmt.Errorf("strategy cannot be nil")
	}
	if minimumWordLength < 0 {
		return nil, fmt.Errorf("minimumWordLength cannot be negative, got: %d", minimumWordLength)
	}

	loadedMap, err := loadEmojiMap()
	if err != nil {
		return nil, fmt.Errorf("failed to load emoji map: %w", err)
	}

	return &Emojifier{
		strategy:          strategy,
		emojiTags:         loadedMap,
		emojiSet:          createEmojiSet(loadedMap),
		minimumWordLength: minimumWordLength,
	}, nil
}

// Emojify applies the configured strategy to add emojis to the given text.
func (e *Emojifier) Emojify(text string) string {
	return e.strategy.Emojify(text, e.minimumWordLength, e.emojiTags, e.emojiSet)
}

// ContainsEmoji returns true if the text contains any emoji characters.
func (e *Emojifier) ContainsEmoji(text string) bool {
	emojis := e.ExtractEmojis(text)
	return len(emojis) > 0
}

// ExtractEmojis returns a slice of all emoji characters found in the text.
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
	// Pre-calculate capacity for better performance
	capacity := 0
	for _, emojis := range emojiTags {
		capacity += len(emojis)
	}

	result := make(map[string]bool, capacity)
	for _, emojis := range emojiTags {
		for _, emoji := range emojis {
			result[emoji] = true
		}
	}
	return result
}

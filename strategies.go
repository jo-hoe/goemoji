package goemoji

import (
	"fmt"
	"strings"
)

const (
	maxKeyLength = 8
)

// EmojifyStrategy defines the interface for different emoji insertion strategies.
type EmojifyStrategy interface {
	Emojify(
		input string,
		minimumWordLength int,
		emojiTags map[string][]string,
		emojiSet map[string]bool,
	) (output string)
}

// ReplaceSubstring replaces words in the text with their corresponding emojis.
type ReplaceSubstring struct{}

// Emojify replaces matching words with emojis in the input text.
func (r ReplaceSubstring) Emojify(
	input string,
	minimumWordLength int,
	emojiTags map[string][]string,
	emojiSet map[string]bool,
) (output string) {
	currentString := strings.ToLower(input)
	for i := maxKeyLength; i > 0; i-- {
		words := strings.Split(currentString, " ")
		tokens := combineTokens(words, i)
		for _, token := range tokens {
			if len(token) < minimumWordLength {
				continue
			}
			emoji, substring := getFirstEmoji(token, emojiTags)
			if substring != "" {
				currentString = strings.Replace(currentString, substring, emoji, 1)
			}
		}
	}
	return currentString
}

// InsertBeforeString inserts emojis before the original text.
type InsertBeforeString struct{}

// Emojify inserts relevant emojis before the input text.
func (i InsertBeforeString) Emojify(
	input string,
	minimumWordLength int,
	emojiTags map[string][]string,
	emojiSet map[string]bool,
) string {
	return fmt.Sprintf("%s %s", getEmojisString(input, minimumWordLength, emojiTags, emojiSet), input)
}

// InsertAfterString inserts emojis after the original text.
type InsertAfterString struct{}

// Emojify inserts relevant emojis after the input text.
func (i InsertAfterString) Emojify(
	input string,
	minimumWordLength int,
	emojiTags map[string][]string,
	emojiSet map[string]bool,
) (output string) {
	return fmt.Sprintf("%s %s", input, getEmojisString(input, minimumWordLength, emojiTags, emojiSet))
}

func getEmojisString(
	input string,
	minimumWordLength int,
	emojiTags map[string][]string,
	emojiSet map[string]bool,
) string {
	emojiString := ReplaceSubstring{}.Emojify(input, minimumWordLength, emojiTags, nil)
	emojies := extractEmojis(emojiString, emojiSet)
	return strings.Join(emojies, "")
}

func extractEmojis(input string, emojiSet map[string]bool) []string {
	if input == "" || len(emojiSet) == 0 {
		return []string{}
	}

	results := make([]string, 0)
	for _, r := range input {
		emoji := string(r)
		if emojiSet[emoji] {
			results = append(results, emoji)
		}
	}

	return results
}

func getFirstEmoji(token string, emojiMap map[string][]string) (emoji, substring string) {
	if emojis, ok := emojiMap[token]; ok {
		return emojis[0], token
	}
	return "", ""
}

func combineTokens(words []string, numWords int) []string {
	if len(words) < numWords || numWords <= 0 {
		return []string{}
	}

	// Pre-allocate slice with known capacity for better performance
	capacity := len(words) - numWords + 1
	tokens := make([]string, 0, capacity)

	for i := 0; i < capacity; i++ {
		token := strings.Join(words[i:i+numWords], " ")
		tokens = append(tokens, token)
	}

	return tokens
}

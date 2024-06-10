package goemoji

import (
	"fmt"
	"strings"
)

const (
	maxKeyLength = 8
)

type EmojifyStrategy interface {
	Emojify(input string, emojiTags map[string][]string, emojiSet map[string]bool) (output string)
}

type ReplaceSubstring struct{}

func (r ReplaceSubstring) Emojify(input string, emojiTags map[string][]string, emojiSet map[string]bool) (output string) {
	currentString := input
	for i := maxKeyLength; i > 0; i-- {
		words := strings.Split(currentString, " ")
		tokens := combineTokens(words, i)
		for _, token := range tokens {
			emoji, substring := getFirstEmoji(token, emojiTags)
			if substring != "" {
				currentString = strings.Replace(currentString, substring, emoji, 1)
			}
		}
	}
	return currentString
}

type InsertBeforeString struct{}

func (i InsertBeforeString) Emojify(input string, emojiTags map[string][]string, emojiSet map[string]bool) string {
	return fmt.Sprintf("%s %s", getEmojisString(input, emojiTags, emojiSet), input)
}

type InsertAfterString struct{}

func (i InsertAfterString) Emojify(input string, emojiTags map[string][]string, emojiSet map[string]bool) (output string) {
	return fmt.Sprintf("%s %s", input, getEmojisString(input, emojiTags, emojiSet))
}

func getEmojisString(input string, emojiTags map[string][]string, emojiSet map[string]bool) string {
	emojiString := ReplaceSubstring{}.Emojify(input, emojiTags, nil)
	emojies := extractEmojis(emojiString, emojiSet)
	return strings.Join(emojies, "")
}

func extractEmojis(input string, emojiSet map[string]bool) []string {
	results := make([]string, 0)

	for _, r := range input {
		if _, ok := emojiSet[string(r)]; ok {
			results = append(results, string(r))
		}
	}

	return results
}

func getFirstEmoji(token string, emojiMap map[string][]string) (emoji string, substring string) {
	if emojis, ok := emojiMap[token]; ok {
		return emojis[0], token
	}
	return "", ""
}

func combineTokens(words []string, numWords int) []string {
	if len(words) < numWords {
		return []string{}
	}

	tokens := make([]string, 0)
	for i := 0; i < len(words)-numWords+1; i++ {
		token := strings.Join(words[i:i+numWords], " ")
		tokens = append(tokens, token)
	}

	return tokens
}

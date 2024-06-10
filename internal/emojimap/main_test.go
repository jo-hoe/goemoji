package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_generateEmojiMap(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		[
			{
				"emoji": "😀",
				"description": "grinning face",
				"category": "Smileys & Emotion",
				"aliases": [
					"grinning_face"
				],
				"tags": [
					"smile",
					"happy"
				],
				"unicode_version": "6.1",
				"ios_version": "6.0"
			},
			{
				"emoji": "🍽️",
				"description": "fork and knife with plate",
				"category": "Food & Drink",
				"aliases": [
					"plate_with_cutlery"
				],
				"tags": [
					"dining",
					"dinner"
				],
				"unicode_version": "7.0",
				"ios_version": "9.1"
			}
		]`))
	}))
	defer ts.Close()

	apiURL := ts.URL

	result := generateEmojiMap(apiURL)
	expectedNumberOfElements := 7
	if len(result) != expectedNumberOfElements {
		t.Errorf("Expected %d emojis, got %d", expectedNumberOfElements, len(result))
	}

	notAllowedCharacter := "_"
	for key := range result {
		if strings.Contains(key, notAllowedCharacter) {
			t.Errorf("key '%s' should not contain '%s'", key, notAllowedCharacter)
		}
	}
}

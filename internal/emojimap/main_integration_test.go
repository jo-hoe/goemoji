package main

import (
	"encoding/json"
	"os"
	"path"
	"reflect"
	"testing"
)

func Test_generateMap(t *testing.T) {
	filePath := path.Join(os.TempDir(), "unittest_emoji_map.json")
	defer os.Remove(filePath)

	generateMap(filePath)

	_, err := os.Stat(filePath)
	if err != nil {
		t.Errorf("Expected output directory to exist, but got error: %v", err)
	}
}

func Test_storeMapToJson_ValidMap(t *testing.T) {
	emojiMap := map[string][]string{
		"smile": {"😄", "😃"},
		"laugh": {"🤣"},
	}
	filePath := path.Join(os.TempDir(), "unittest_emoji_map.json")
	defer os.Remove(filePath)

	storeMapToJson(emojiMap, filePath)

	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	var actualMap map[string][]string
	err = json.Unmarshal(data, &actualMap)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %v", err)
	}

	if !reflect.DeepEqual(emojiMap, actualMap) {
		t.Errorf("Expected emoji map: %v, got: %v", emojiMap, actualMap)
	}
}

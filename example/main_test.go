package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func Test_example(t *testing.T) {
	expectedLogs := []string{"🎶 puts a 😄 on my face.\n", "🎶😄 Music puts a smile on my face."}
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	example()

	got := buf.String()

	for _, expectedLog := range expectedLogs {
		if !strings.Contains(got, expectedLog) {
			t.Errorf("expected substring '%s', but got '%s'", expectedLog, got)
		}
	}
}

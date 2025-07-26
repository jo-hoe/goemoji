# Go Emoji 🎉

[![Test Status](https://github.com/jo-hoe/goemoji/workflows/test/badge.svg)](https://github.com/jo-hoe/goemoji/actions?workflow=test)
[![Lint Status](https://github.com/jo-hoe/goemoji/workflows/lint/badge.svg)](https://github.com/jo-hoe/goemoji/actions?workflow=lint)
[![Go Report Card](https://goreportcard.com/badge/github.com/jo-hoe/goemoji)](https://goreportcard.com/report/github.com/jo-hoe/goemoji)
[![Coverage Status](https://coveralls.io/repos/github/jo-hoe/goemoji/badge.svg?branch=main)](https://coveralls.io/github/jo-hoe/goemoji?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/jo-hoe/goemoji.svg)](https://pkg.go.dev/github.com/jo-hoe/goemoji)

A Go library that adds emojis to text using configurable strategies. Transform plain text into expressive, emoji-rich content with multiple insertion modes and customizable word matching.

## Features

- 🎯 **Multiple Strategies**: Replace words, insert before/after text
- 🔧 **Configurable**: Set minimum word length for matching
- 🚀 **Thread-Safe**: Safe for concurrent use
- 📦 **Zero Dependencies**: Pure Go implementation
- 🎨 **Rich Emoji Database**: Comprehensive emoji-to-word mapping

## Installation

```bash
go get github.com/jo-hoe/goemoji
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/jo-hoe/goemoji"
)

func main() {
    // Create emojifier with default settings
    emojifier, err := goemoji.NewDefaultEmojifier()
    if err != nil {
        log.Fatal(err)
    }
    
    text := "Music puts a smile on my face."
    result := emojifier.Emojify(text)
    fmt.Println(result) // Output: 🎶 puts a 😄 on my face.
}
```

## Strategies

The library supports three different emoji insertion strategies:

### 1. ReplaceSubstring (Default)
Replaces matching words with their corresponding emojis:

```go
emojifier, _ := goemoji.NewDefaultEmojifier()
result := emojifier.Emojify("I love music and dancing")
// Output: "I love 🎶 and 💃"
```

### 2. InsertBeforeString
Inserts relevant emojis before the original text:

```go
emojifier, _ := goemoji.NewEmojifier(goemoji.InsertBeforeString{}, 4)
result := emojifier.Emojify("Music puts a smile on my face")
// Output: "🎶😄 Music puts a smile on my face"
```

### 3. InsertAfterString
Inserts relevant emojis after the original text:

```go
emojifier, _ := goemoji.NewEmojifier(goemoji.InsertAfterString{}, 4)
result := emojifier.Emojify("Music puts a smile on my face")
// Output: "Music puts a smile on my face 🎶😄"
```

## Advanced Usage

### Custom Minimum Word Length
Control which words get matched by setting a minimum length:

```go
// Only match words with 6+ characters
emojifier, _ := goemoji.NewEmojifier(goemoji.ReplaceSubstring{}, 6)
result := emojifier.Emojify("I love music")
// "music" (5 chars) won't be replaced, but longer words will
```

### Emoji Detection
Check if text contains emojis or extract them:

```go
emojifier, _ := goemoji.NewDefaultEmojifier()

// Check if text contains emojis
hasEmojis := emojifier.ContainsEmoji("Hello 👋 world")
fmt.Println(hasEmojis) // true

// Extract all emojis from text
emojis := emojifier.ExtractEmojis("Music 🎶 and dance 💃")
fmt.Println(emojis) // ["🎶", "💃"]
```

## Documentation

For complete API documentation, examples, and detailed usage instructions, visit the [Go Reference](https://pkg.go.dev/github.com/jo-hoe/goemoji).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

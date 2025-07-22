package main

import (
	"log"

	"github.com/jo-hoe/goemoji"
)

func main() {
	example()
}

func example() {
	const minWordLength = 4
	input := "Music puts a smile on my face."

	emojifier, err := goemoji.NewDefaultEmojifier()
	if err != nil {
		panic(err)
	}
	log.Println(emojifier.Emojify(input))

	emojifier, err = goemoji.NewEmojifier(goemoji.InsertBeforeString{}, minWordLength)
	if err != nil {
		panic(err)
	}
	log.Println(emojifier.Emojify(input))
}

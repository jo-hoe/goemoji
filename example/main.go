package main

import "github.com/jo-hoe/goemoji"

func main() {
	example()
}

func example() {
	input := "Music puts a smile on my face."

	emojifier, err := goemoji.NewDefaultEmojifier()
	if err != nil {
		panic(err)
	}
	println(emojifier.Emojify(input))

	emojifier, err = goemoji.NewEmojifier(goemoji.InsertBeforeString{}, 4)
	if err != nil {
		panic(err)
	}
	println(emojifier.Emojify(input))
}

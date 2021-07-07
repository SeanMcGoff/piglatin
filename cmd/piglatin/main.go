package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SeanMcGoff/piglatin"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"
const vowels = "aeiou"

func main() {
	var text string
	var help bool

	flag.StringVar(&text, "t", "", "Input Text to Translate")
	flag.BoolVar(&help, "help", false, "Display Usage Information")

	flag.Parse()

	if help || text == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	translation, err := piglatin.ToPigLatin(text)
	if err != nil {
		exitWithError(err)
	}
	fmt.Println(translation)
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}

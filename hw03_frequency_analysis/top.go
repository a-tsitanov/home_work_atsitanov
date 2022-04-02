package hw03frequencyanalysis

import (
	"regexp"
)

var wordsRegular = regexp.MustCompile(`\s+`)

func splitText(text string) []string {
	return wordsRegular.Split(text, -1)
}

func Top10(input string) []string {
	words := splitText(input)
	wordsCount := NewWordsText()
	for _, word := range words {
		if word == "" {
			continue
		}
		_ = wordsCount.addWord(word)
	}
	return wordsCount.getCountWords(10)
}

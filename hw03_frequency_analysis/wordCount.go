package hw03frequencyanalysis

import "sort"

type Word struct {
	word  string
	count int
}

func NewWord(word string) *Word {
	return &Word{
		word:  word,
		count: 1,
	}
}

func (w *Word) increment() {
	w.count++
}

type WordsText struct {
	Words map[string]*Word
}

func NewWordsText() *WordsText {
	return &WordsText{
		Words: make(map[string]*Word),
	}
}

func (wText *WordsText) getWords() []*Word {
	words := make([]*Word, 0, len(wText.Words))
	for _, val := range wText.Words {
		words = append(
			words,
			val,
		)
	}
	sort.Slice(words, func(i, j int) bool {
		if words[i].count == words[j].count {
			return words[i].word < words[j].word
		}
		return words[i].count > words[j].count
	})

	return words
}

func (wText *WordsText) getCountWords(count int) []string {
	words := wText.getWords()
	lenWords := len(words)
	if lenWords > count {
		words = words[:10]
	}
	result := make([]string, 0)
	for _, word := range words {
		result = append(result, word.word)
	}
	return result
}

func (wText *WordsText) addWord(word string) bool {
	if _, ok := wText.Words[word]; ok {
		wText.Words[word].increment()
		return false
	}
	wText.Words[word] = NewWord(word)
	return true
}

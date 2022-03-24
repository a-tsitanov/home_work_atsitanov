package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	result := strings.Builder{}
	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		currentSymbol := runes[i]
		if i == 0 && unicode.IsDigit(currentSymbol) {
			return "", ErrInvalidString
		}
		if i == len(runes)-1 {
			if !unicode.IsDigit(currentSymbol) {
				result.WriteRune(currentSymbol)
			}
			break
		}

		nextSymbol := runes[i+1]

		if unicode.IsDigit(currentSymbol) && unicode.IsDigit(nextSymbol) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(nextSymbol) {
			countRepeat, err := strconv.Atoi(string(nextSymbol))
			if err == nil {
				result.WriteString(strings.Repeat(string(currentSymbol), countRepeat))
			}
			continue
		}

		if unicode.IsDigit(currentSymbol) {
			continue
		} else {
			result.WriteRune(currentSymbol)
		}
	}
	return result.String(), nil
}

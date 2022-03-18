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
	lenStr := len(input)

	for i, symb := range input {
		if i == 0 && unicode.IsDigit(symb) {
			return "", ErrInvalidString
		}

		if i == lenStr-1 {
			if !unicode.IsDigit(symb) {
				result.WriteRune(symb)
			}
			break
		}

		nextSymb := rune(input[i+1])

		if unicode.IsDigit(symb) && unicode.IsDigit(nextSymb) {
			return "", ErrInvalidString
		}

		if (unicode.IsLetter(symb) || unicode.IsSpace(symb)) && (unicode.IsLetter(nextSymb) || unicode.IsSpace(nextSymb)) {
			result.WriteRune(symb)
		}

		if (unicode.IsLetter(symb) || unicode.IsSpace(symb)) && unicode.IsDigit(nextSymb) {
			countRepeat, err := strconv.Atoi(string(input[i+1]))
			if err == nil {
				result.WriteString(strings.Repeat(string(symb), countRepeat))
			}
		}
	}
	return result.String(), nil
}

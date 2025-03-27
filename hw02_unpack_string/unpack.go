package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	var builder strings.Builder
	runes := []rune(packedString)

	for i := 1; i <= len(runes)-1; i++ {
		prevSymbol := runes[i-1]
		curSymbol := runes[i]

		if isDigit(curSymbol) && !isDigit(prevSymbol) {
			num, _ := strconv.Atoi(string(curSymbol))
			builder.WriteString(strings.Repeat(string(prevSymbol), num))
		}

		if !isDigit(curSymbol) && !isDigit(prevSymbol) {
			builder.WriteString(string(prevSymbol))
		}

		if !isDigit(curSymbol) && i == len(runes)-1 {
			builder.WriteString(string(curSymbol))
		}

		if isDigit(prevSymbol) && i == 1 {
			return "", ErrInvalidString
		}
		if isDigit(prevSymbol) && isDigit(curSymbol) {
			return "", ErrInvalidString
		}
	}

	return builder.String(), nil
}

func isDigit(symbol rune) bool {
	return unicode.IsDigit(symbol)
}

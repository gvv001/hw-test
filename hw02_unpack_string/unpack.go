package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	var builder strings.Builder
	runes := []rune(packedString)

	for i := 1; i <= len(runes)-1; i++ {
		prevSymbol := string(runes[i-1])
		curSymbol := string(runes[i])

		if count, err := strconv.Atoi(curSymbol); err == nil {
			if _, err := strconv.Atoi(prevSymbol); err != nil {
				builder.WriteString(strings.Repeat(prevSymbol, count))
			}
		}

		if _, err := strconv.Atoi(curSymbol); err != nil {
			if _, err := strconv.Atoi(prevSymbol); err != nil {
				builder.WriteString(prevSymbol)
			}
		}

		if _, err := strconv.Atoi(curSymbol); err != nil && i == len(runes)-1 {
			builder.WriteString(curSymbol)
		}

		if _, err := strconv.Atoi(prevSymbol); i == 1 && err == nil {
			return "", ErrInvalidString
		}

		if _, err := strconv.Atoi(prevSymbol); err == nil {
			if _, err := strconv.Atoi(curSymbol); err == nil {
				return "", ErrInvalidString
			}
		}
	}

	return builder.String(), nil
}

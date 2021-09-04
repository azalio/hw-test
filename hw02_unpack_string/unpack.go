package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func getRepeatedString(r rune, d int) (string, error) {
	if unicode.IsDigit(r) {
		return "", ErrInvalidString
	}

	return strings.Repeat(string(r), d), nil
}

func Unpack(s string) (string, error) {
	r, _ := utf8.DecodeRuneInString(s)
	if unicode.IsDigit(r) {
		return "", ErrInvalidString
	}

	if s == "" {
		return "", nil
	}

	builder := strings.Builder{}
	accumulatedString := strings.Builder{}
	var prevRune rune
	var digitFounded bool
	var digit, size int

	for len(s) > 0 {
		r, size = utf8.DecodeRuneInString(s) // get one rune from string

		if unicode.IsDigit(r) {
			digit = int(r) - 48 // get digit from rune
			digitFounded = true
		}

		if digitFounded {
			repeatedString, err := getRepeatedString(prevRune, digit)
			if err != nil {
				return "", err
			}
			builder.WriteString(repeatedString)
			digitFounded = false
		} else if (prevRune < 48 || prevRune > 57) && prevRune != 0 {
			// not digit and not default for rune
			builder.WriteRune(prevRune)
		}

		accumulatedString.WriteString(builder.String())
		builder.Reset()

		prevRune = r
		s = s[size:]
	}

	if !unicode.IsDigit(prevRune) { // add last rune from string
		accumulatedString.WriteRune(r)
	}

	return accumulatedString.String(), nil
}

package service

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func StringConverter(s string) (string, error) {

	if s == "" {

		return "", errors.New("Empty line")
	}

	s = strings.TrimSpace(s)
	if s == "" {

		return "", errors.New("Empty line")
	}

	var MorseCharsCounter, TextCharsCounter int

	for _, ch := range s {

		morseChars := ".-/·"
		if strings.ContainsRune(morseChars, ch) {
			MorseCharsCounter++
			continue
		}

		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			TextCharsCounter++
		}
	}

	IsMorse := false
	if TextCharsCounter == 0 && MorseCharsCounter > 0 {
		IsMorse = true
	} else if MorseCharsCounter > TextCharsCounter {
		IsMorse = true
	}

	if IsMorse {
		Normalized := strings.ReplaceAll(s, "\r\n", " ")
		Normalized = strings.ReplaceAll(Normalized, "\n", " ")
		Normalized = strings.TrimSpace(Normalized)
		return morse.ToText(Normalized), nil
	}

	NormalizedText := strings.Join(strings.Fields(s), " ")
	return morse.ToMorse(NormalizedText), nil
}

func Converter(r io.Reader) (string, error) {

	if r == nil {
		return " ", errors.New("Empty line")
	}

	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)

	if err != nil {
		return " ", err
	}
	content := buf.String()
	return StringConverter(content)
}

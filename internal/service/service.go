package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// AutoConvert получает строку и определяет её тип (текст или код Морзе)
func AutoConvert(input string) (string, error) {
	isMorseCode := true
	for _, j := range input {
		if !strings.ContainsRune(".-/", j) {
			isMorseCode = false
			break
		}
	}

	if isMorseCode {
		result := morse.ToText(input)
		return result, nil
	} else {
		result := morse.ToMorse(input)
		return result, nil
	}
}

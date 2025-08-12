package service

import (
	//"strings"

	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// AutoConvert получает строку и определяет её тип (текст или код Морзе)
func AutoConvert(input string) (string, error) {
	trimmedInput := strings.TrimSpace(input)
	upperCaseInput := strings.ToUpper(trimmedInput)

	isMorseCode := true
	for _, j := range upperCaseInput {
		if !strings.ContainsRune(".-/ ", j) {
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

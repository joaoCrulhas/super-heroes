package encrypter

import (
	"unicode"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
)

type EncryptValidators func(string) error

func ValidateEmptyInput(value string) error {
	if value == "" {
		return domain.EmptyString() // Fix the function call
	}
	return nil
}

func ValidateSpecialCharacters(value string) error {
	for _, r := range value {
		if !unicode.IsLetter(r) {
			return domain.InvalidCharacters()
		}
	}
	return nil
}

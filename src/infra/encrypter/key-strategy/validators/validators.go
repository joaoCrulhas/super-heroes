package encrypter

import (
	"unicode"

	custom_errors "github.com/joaoCrulhas/omnevo-super-heroes/src/error"
)

type EncryptValidators func(string) error

func ValidateEmptyInput(value string) error {
	if value == "" {
		return custom_errors.EmptyString()
	}
	return nil
}

func ValidateSpecialCharacters(value string) error {
	for _, r := range value {
		if !unicode.IsLetter(r) {
			return custom_errors.InvalidCharacters()
		}
	}
	return nil
}

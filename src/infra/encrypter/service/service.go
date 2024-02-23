package encrypter

import (
	"unicode"

	custom_errors "github.com/joaoCrulhas/omnevo-super-heroes/src/error"
)

type Service struct {
	key uint32
}

func NewEncryptService(key uint32) Service {
	return Service{
		key: key,
	}
}

func (service *Service) hasInvalidChars(value string) bool {
	for _, r := range value {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (service *Service) Encrypt(value string) (string, error) {
	if value == "" {
		return "", custom_errors.EmptyString()
	}
	if !service.hasInvalidChars(value) {
		return "", custom_errors.InvalidCharacters()
	}
	return value, nil
}

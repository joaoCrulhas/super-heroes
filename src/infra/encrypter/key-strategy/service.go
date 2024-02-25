package encrypter

import (
	"unicode"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	validators "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/key-strategy/validators"
)

type Service struct {
	key        uint32
	dictionary domain.Dictionary
	validators []validators.EncryptValidators
}

func NewEncryptService(key uint32, dictionary domain.Dictionary, fnValidators ...validators.EncryptValidators) *Service {
	return &Service{
		key:        key,
		dictionary: dictionary,
		validators: fnValidators,
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
	var encryptedValue string
	for _, fn := range service.validators {
		if err := fn(value); err != nil {
			return "", err
		}
	}
	for _, letter := range value {
		key := service.dictionary.GetKey(letter)
		var t rune
		if uint32(key)+service.key > 26 {
			var value = (uint32(key) + service.key) % 26
			t = service.dictionary.GetValue(int(value))
		} else {
			var value = uint32(key) + service.key
			t = service.dictionary.GetValue(int(value))
		}
		encryptedValue += string(t)
	}
	return encryptedValue, nil
}

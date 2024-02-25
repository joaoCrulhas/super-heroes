package encrypter

import (
	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	validators "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/key-strategy/validators"
)

type Service struct {
	key        int
	dictionary domain.Dictionary
	validators []validators.EncryptValidators
}

func NewEncryptService(key int, dictionary domain.Dictionary, fnValidators ...validators.EncryptValidators) *Service {
	return &Service{
		key:        key,
		dictionary: dictionary,
		validators: fnValidators,
	}
}

func (service *Service) Encrypt(input string) (string, error) {
	var encryptedValue string
	for _, fn := range service.validators {
		if err := fn(input); err != nil {
			return "", err
		}
	}
	alphabetLength := service.dictionary.GetAlphabetLength()
	var t rune
	for _, letter := range input {
		key := service.dictionary.GetKey(letter)
		var value int
		if (key)+service.key > alphabetLength {
			value = ((key) + service.key) % alphabetLength
		} else {
			value = (key) + service.key
		}
		t = service.dictionary.GetValue(value)
		encryptedValue += string(t)
	}
	return encryptedValue, nil
}

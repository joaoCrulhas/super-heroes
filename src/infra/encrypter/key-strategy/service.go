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
	err := service.execValidators(input)
	if err != nil {
		return "", err
	}

	var encryptedValue string
	var t rune
	var value int
	alphabetLength := service.dictionary.GetAlphabetLength()
	for _, letter := range input {
		key := service.dictionary.GetKey(letter)
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

func (service *Service) execValidators(input string) error {
	for _, fn := range service.validators {
		if err := fn(input); err != nil {
			return err
		}
	}
	return nil
}

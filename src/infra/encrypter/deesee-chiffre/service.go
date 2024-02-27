package encrypter

import (
	"strings"

	validators "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/deesee-chiffre/validators"
)



type EncryptDeeSeeChiffreService struct {
	key           rune
	validators    []validators.EncryptValidators
	maxShiftValue rune
	minShiftValue rune
}

func NewEncryptDeeSeeChiffreService(key rune, minShiftValue rune, maxShiftValue rune, fnValidators ...validators.EncryptValidators) *EncryptDeeSeeChiffreService {
	return &EncryptDeeSeeChiffreService{
		key:           key,
		validators:    fnValidators,
		maxShiftValue: maxShiftValue,
		minShiftValue: minShiftValue,
	}
}

func (service *EncryptDeeSeeChiffreService) Encrypt(input string) (string, error) {
	err := service.execValidators(input)
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	for _, letter := range input {
		key := letter + service.key
		if key > service.maxShiftValue {
			key = key%service.maxShiftValue + service.minShiftValue
		}
		builder.WriteRune(rune(key))
	}
	return strings.ToLower(builder.String()), nil
}

func (service *EncryptDeeSeeChiffreService) execValidators(input string) error {
	for _, fn := range service.validators {
		if err := fn(input); err != nil {
			return err
		}
	}
	return nil
}

package encrypter

import (
	"unicode"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	custom_errors "github.com/joaoCrulhas/omnevo-super-heroes/src/error"
)

type Service struct {
	key        uint32
	dictionary domain.Dictionary
}

func NewEncryptService(key uint32, dictionary domain.Dictionary) Service {
	return Service{
		key:        key,
		dictionary: dictionary,
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
	if value == "" {
		return "", custom_errors.EmptyString()
	}
	if !service.hasInvalidChars(value) {
		return "", custom_errors.InvalidCharacters()
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

// Ao receber uma letra eu preciso fazer o seguinte, acessar o dicionário e pegar a chave do map, ao pegar a chave do map tenho que fazer a seguitne conta
// currentCharCode + key > 26 ? ( currentCharCode + key) % 26 : currentCharCode + key

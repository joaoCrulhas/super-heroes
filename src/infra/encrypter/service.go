package encrypter

import custom_errors "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/error"

type Service struct{}

func NewEncryptService() Service {
	return Service{}
}

func (e *Service) Encrypt(value string) (string, error) {
	if value == "" {
		return "", custom_errors.EmptyString()
	}
	return value, nil
}

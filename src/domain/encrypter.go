package domain

type (
	Encrypter interface {
		Encrypt(input string) (string, error)
	}
)

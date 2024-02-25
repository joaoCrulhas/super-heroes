package domain

type (
	Encrypt interface {
		Encrypt(input string) (string, error)
	}
)

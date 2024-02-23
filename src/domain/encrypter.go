package domain

type (
	Encrypt interface {
		Encrypt(value string) (string, error)
	}
)

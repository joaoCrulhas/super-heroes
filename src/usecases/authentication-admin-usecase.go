package usecases

import (
	errors "github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
)

type AuthenticationAdmin struct {
}

func NewAuthenticationAdmin() *AuthenticationAdmin {
	return &AuthenticationAdmin{}
}

func (a *AuthenticationAdmin) Auth(headers map[string][]string) (bool, error) {
	if _, ok := headers["X-Dee-See-Admin-Key"]; !ok {
		return false, errors.Unauthorized("Unauthorized")
	}

	if headers["X-Dee-See-Admin-Key"][0] != "myadminkey" {
		return false, errors.Unauthorized("Unauthorized")
	}

	return true, nil
}

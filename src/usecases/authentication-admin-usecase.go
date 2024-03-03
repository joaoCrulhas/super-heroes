package usecases

import (
	errors "github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
)

type AuthenticationAdmin struct {
	apiKey string
}

func NewAuthenticationAdmin() *AuthenticationAdmin {
	return &AuthenticationAdmin{}
}

func (c *AuthenticationAdmin) Inject(
	cfg *struct {
		APIKey string `inject:"config:deesee.apikey"`
	},
) *AuthenticationAdmin {
	if cfg != nil {
		c.apiKey = cfg.APIKey
	}
	return c
}

func (a *AuthenticationAdmin) Auth(headers map[string][]string) (bool, error) {
	if _, ok := headers["X-Dee-See-Admin-Key"]; !ok {
		return false, errors.Unauthorized("Unauthorized")
	}

	if headers["X-Dee-See-Admin-Key"][0] != a.apiKey {
		return false, errors.Unauthorized("Unauthorized")
	}

	return true, nil
}

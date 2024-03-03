package presentation

import (
	"fmt"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"golang.org/x/exp/slices"
)

// create a func as type
type Validator func(string) error

func ValidateSuperPowerInput(value string) error {
	isValid := slices.Contains(domain.SuperPowers, value)
	if !isValid {
		errMessage := fmt.Sprintf("superpower %v not found", value)
		return domain.BadRequest(errMessage)
	}
	return nil
}

package domain

import "context"

var SuperPowers = []string{"strength", "speed", "flight", "invulnerability", "healing"}

type (
	Superhero struct {
		Name        string   `json:"name"`
		Identity    Identity `json:"identity"`
		Birthday    string   `json:"birthday"`
		Superpowers []string `json:"superpowers"`
	}
	Identity struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	// This is the usecases for the SuperHero Domain
	SuperHeroUseCase interface {
		Fetch(c context.Context) ([]Superhero, error)
		GetBySuperPower(c context.Context, powers map[string]any) ([]Superhero, error)
	}
)

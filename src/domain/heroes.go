package domain

import (
	"context"
)

var SuperPowers = []string{"strength", "speed", "flight", "invulnerability", "healing"}

type (
	Superhero struct {
		Name        string   `json:"name,omitempty"`
		Identity    Identity `json:"identity,omitempty"`
		Birthday    string   `json:"birthday,omitempty"`
		Superpowers []string `json:"superpowers"`
	}
	Identity struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}
	// This is the usecases for the SuperHero Domain
	SuperHeroUseCase interface {
		Fetch(ctx context.Context, filter map[string][]string) ([]Superhero, error)
		GetBySuperPower(ctx context.Context, powers []string) ([]Superhero, error)
		EncryptIdentity(ctx context.Context, identity Identity) (string, error)
		Create(ctx context.Context, superHero Superhero) (SuperHeroWithEncryptIdentity, error)
	}

	SuperHeroWithEncryptIdentity struct {
		Name        string   `json:"name,omitempty"`
		Identity    string   `json:"identity,omitempty"`
		Birthday    string   `json:"birthday,omitempty"`
		Superpowers []string `json:"superpowers"`
	}
)

func ParseSuperHero(superHero Superhero) SuperHeroWithEncryptIdentity {
	return SuperHeroWithEncryptIdentity{
		Name:        superHero.Name,
		Identity:    superHero.Identity.FirstName + " " + superHero.Identity.LastName,
		Birthday:    superHero.Birthday,
		Superpowers: superHero.Superpowers,
	}
}
func ParseResponse(superHeroes []Superhero) []SuperHeroWithEncryptIdentity {
	var response []SuperHeroWithEncryptIdentity
	for _, superHero := range superHeroes {
		parsedSuperHero := ParseSuperHero(superHero)
		response = append(response, parsedSuperHero)
	}
	return response

}
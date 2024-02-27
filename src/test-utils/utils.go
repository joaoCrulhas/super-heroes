package testutils

import "github.com/joaoCrulhas/omnevo-super-heroes/src/domain"

func GetSuperHeroes() []domain.Superhero {
	heroes := []domain.Superhero{ // Change from array to slice
		{
			Name: "superHero1",
			Identity: domain.Identity{
				FirstName: "Snyder",
				LastName:  "Johnston",
			},
			Birthday:    "1990-04-14",
			Superpowers: []string{"flight", "strength", "invulnerability"},
		},
		{
			Name: "Super Hero 2",
			Identity: domain.Identity{
				FirstName: "Petra",
				LastName:  "Sharpe",
			},
			Birthday:    "1973-04-18", // Batman's first appearance in comics
			Superpowers: []string{},
		},
	}
	return heroes
}

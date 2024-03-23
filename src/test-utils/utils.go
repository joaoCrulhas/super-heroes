package testutils

import "github.com/joaoCrulhas/omnevo-super-heroes/src/domain"

func GetSuperHeroes() domain.SuperHerosData {
	heroes := domain.SuperHerosData{}
	heroes[1] = &domain.Superhero{
		ID:   1,
		Name: "superHero1",
		Identity: domain.Identity{
			FirstName: "Snyder",
			LastName:  "Johnston",
		},
		Birthday:    "1990-04-14",
		Superpowers: []string{"flight", "strength", "invulnerability"},
	}
	heroes[2] = &domain.Superhero{
		ID:   2,
		Name: "superHero2",
		Identity: domain.Identity{
			FirstName: "Test2",
			LastName:  "TestLastName2",
		},
		Birthday:    "1990-04-14",
		Superpowers: []string{},
	}
	heroes[3] = &domain.Superhero{
		ID:   3,
		Name: "superHero3",
		Identity: domain.Identity{
			FirstName: "Test3",
			LastName:  "TestLastName3",
		},
		Birthday:    "1990-04-14",
		Superpowers: []string{"healing"},
	}
	return heroes
}

package db

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"golang.org/x/exp/slices"
)

// If provided with a nil slice, it will read the file from the path provided
func NewSuperHeroMemoryRepository(superHeroes domain.SuperHerosData) (*SuperHeroMemoryRepository, error) {
	if len(superHeroes) == 0 {
		var err error
		superHeroes, err = ReadSuperHeroFile("superheroes.json")
		if err != nil {
			return nil, err
		}
	}
	return &SuperHeroMemoryRepository{
		superHeroes: superHeroes,
	}, nil
}

type SuperHeroMemoryRepository struct {
	superHeroes domain.SuperHerosData
}

func (r *SuperHeroMemoryRepository) Fetch(c context.Context) (domain.SuperHerosData, error) {
	return r.superHeroes, nil
}

func (r *SuperHeroMemoryRepository) FindByFilter(c context.Context, filter map[string][]string) (domain.SuperHerosData, error) {
	heroes := domain.SuperHerosData{}
	for _, hero := range r.superHeroes {
		for k, v := range filter {
			if k == "superpowers" {
				for _, s := range v {
					if slices.Contains(hero.Superpowers, s) {
						heroes[hero.ID] = hero
					}
				}
			}
		}
	}
	return heroes, nil
}

func (r *SuperHeroMemoryRepository) Create(c context.Context, item *domain.Superhero) (*domain.Superhero, error) {
	r.superHeroes[len(r.superHeroes)+1] = item
	return item, nil
}

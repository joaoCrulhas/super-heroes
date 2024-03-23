package db

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"golang.org/x/exp/slices"
)

// If provided with a nil slice, it will read the file from the path provided
func NewSuperHeroMemoryRepository(superHeroes map[int]domain.Superhero) (*SuperHeroMemoryRepository, error) {
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
	superHeroes map[int]domain.Superhero
}

func (r *SuperHeroMemoryRepository) Fetch(c context.Context) (map[int]domain.Superhero, error) {
	return r.superHeroes, nil
}

func (r *SuperHeroMemoryRepository) FindByFilter(c context.Context, filter map[string][]string) (map[int]domain.Superhero, error) {
	heroes := map[int]domain.Superhero{}
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

func (r *SuperHeroMemoryRepository) Create(c context.Context, item domain.Superhero) (domain.Superhero, error) {
	r.superHeroes[len(r.superHeroes)+1] = item
	return item, nil
}

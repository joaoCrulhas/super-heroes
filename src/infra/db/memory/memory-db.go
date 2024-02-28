package db

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"golang.org/x/exp/slices"
)

// If provided with a nil slice, it will read the file from the path provided
func NewSuperHeroMemoryRepository(superHeroes []domain.Superhero) *SuperHeroMemoryRepository {
	if len(superHeroes) == 0 {
		var err error
		superHeroes, err = ReadSuperHeroFile("superheroes.json")
		if err != nil {
			// Create zap logger
			panic(err)
		}
	}
	return &SuperHeroMemoryRepository{
		superHeroes: superHeroes,
	}
}

type SuperHeroMemoryRepository struct {
	superHeroes []domain.Superhero
}

func (r *SuperHeroMemoryRepository) Fetch(c context.Context) ([]domain.Superhero, error) {
	return r.superHeroes, nil
}

func (r *SuperHeroMemoryRepository) FindByFilter(c context.Context, filter map[string]any) ([]domain.Superhero, error) {
	var heroes []domain.Superhero
	for _, hero := range r.superHeroes {
		for k, v := range filter {
			if k == "superpowers" {
				cast, ok := v.(string)
				if ok && slices.Contains(hero.Superpowers, cast) {
					heroes = append(heroes, hero)
				}

			}
		}
	}
	return heroes, nil
}

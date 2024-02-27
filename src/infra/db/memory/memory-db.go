package db

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"golang.org/x/exp/slices"
)

func NewSuperHeroMemoryRepository(superHeroes []domain.Superhero) SuperHeroMemoryRepository {
	return SuperHeroMemoryRepository{
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
				if !ok {
					continue
				}
				if hero.Superpowers == nil || len(hero.Superpowers) == 0 {
					continue
				}
				if slices.Contains(hero.Superpowers, cast) {
					heroes = append(heroes, hero)
				}

			}
		}
	}
	return heroes, nil
}

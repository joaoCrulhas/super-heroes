package db

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
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

func (r *SuperHeroMemoryRepository) GetBySuperPower(c context.Context, powers []string) ([]domain.Superhero, error) {
	var heroes []domain.Superhero
	for _, hero := range r.superHeroes {
		for _, power := range hero.Superpowers {
			for _, p := range powers {
				if power == p {
					heroes = append(heroes, hero)
				}
			}
		}
	}
	return heroes, nil
}

package db

import (
	"context"
	"sync"

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
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, hero := range r.superHeroes {
		wg.Add(1)
		go func(h domain.Superhero) {
			defer wg.Done()
			for k, v := range filter {
				if k == "superpowers" {
					for _, s := range v {
						if slices.Contains(h.Superpowers, s) {
							mu.Lock()
							heroes[h.ID] = &h
							mu.Unlock()
						}
					}
				}
			}
		}(*hero)
	}
	wg.Wait()
	return heroes, nil
}

func (r *SuperHeroMemoryRepository) Create(c context.Context, item *domain.Superhero) (*domain.Superhero, error) {
	r.superHeroes[len(r.superHeroes)+1] = item
	return item, nil
}

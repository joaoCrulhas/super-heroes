package usecases

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/infra/db"
)

type superHeroUseCase struct {
	repository db.Repository[domain.Superhero]
}

func NewSuperHeroUseCase(repository db.Repository[domain.Superhero]) *superHeroUseCase {
	return &superHeroUseCase{repository: repository}
}

func (su *superHeroUseCase) Inject(repository db.Repository[domain.Superhero]) {
	su.repository = repository
}

func (su *superHeroUseCase) Fetch(ctx context.Context) ([]domain.Superhero, error) {
	return su.repository.Fetch(ctx)
}

func (su *superHeroUseCase) GetBySuperPower(ctx context.Context, powers map[string]any) ([]domain.Superhero, error) {
	return su.repository.FindByFilter(ctx, powers)
}

package presentation

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	presentation "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
)

type FetchSuperHeroController struct {
	superHerUseCase domain.SuperHeroUseCase
}

func NewFetchController(superHerUseCase domain.SuperHeroUseCase) *FetchSuperHeroController {
	return &FetchSuperHeroController{
		superHerUseCase: superHerUseCase,
	}
}

func (controller *FetchSuperHeroController) Handle(ctx context.Context, request presentation.Request[any]) presentation.Response[[]domain.Superhero] {
	data, err := controller.superHerUseCase.Fetch(ctx)
	if err != nil {
		return presentation.Response[[]domain.Superhero]{
			StatusCode: 500,
			Data:       nil,
		}
	}
	return presentation.Response[[]domain.Superhero]{
		StatusCode: 200,
		Data:       data,
	}
}

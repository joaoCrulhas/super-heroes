package presentation

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	presentation "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
)

type FilterSuperHeroController struct {
	superHerUseCase domain.SuperHeroUseCase
}

func NewFilterController(superHerUseCase domain.SuperHeroUseCase) *FilterSuperHeroController {
	return &FilterSuperHeroController{
		superHerUseCase: superHerUseCase,
	}
}

func (controller *FilterSuperHeroController) Handle(ctx context.Context, request presentation.Request[map[string]any]) presentation.Response[[]domain.Superhero] {
	data, err := controller.superHerUseCase.GetBySuperPower(ctx, request.Body)
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

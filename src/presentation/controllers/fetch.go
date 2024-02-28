package presentation

import (
	"context"
	"errors"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	presentation "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
	presentation_adapter "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation/controllers/adapter"
)

type FetchSuperHeroController struct {
	superHerUseCase domain.SuperHeroUseCase
	responder       *web.Responder
}

func (controller *FetchSuperHeroController) Inject(responder *web.Responder, superHerUseCase domain.SuperHeroUseCase) {
	controller.responder = responder
	controller.superHerUseCase = superHerUseCase
}

func NewFetchController(superHerUseCase domain.SuperHeroUseCase) *FetchSuperHeroController {
	return &FetchSuperHeroController{
		superHerUseCase: superHerUseCase,
	}
}

func (controller *FetchSuperHeroController) Wrapper(ctx context.Context, r *web.Request) web.Result {
	request, err := presentation_adapter.AdapterRequest[any](r)
	if err != nil {
		return controller.responder.ServerError(domain.BadRequest())
	}
	response := controller.Handle(ctx, request)
	if response.Error != nil {
		return controller.responder.ServerError(errors.New("error"))
	}
	return controller.responder.Data(response.Data).Status(response.StatusCode)
}

func (controller *FetchSuperHeroController) Handle(ctx context.Context, request presentation.Request[any]) presentation.Response[[]domain.Superhero] {
	data, err := controller.superHerUseCase.Fetch(ctx)
	if err != nil {
		return presentation.Response[[]domain.Superhero]{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}
	return presentation.Response[[]domain.Superhero]{
		StatusCode: http.StatusOK,
		Data:       data,
	}
}

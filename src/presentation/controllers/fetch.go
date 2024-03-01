package presentation

import (
	"context"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	presentation "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
	presentation_adapter "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation/controllers/adapter"
)

type SuperHeroResponse []domain.SuperHeroWithEncryptIdentity
type FetchSuperHeroController struct {
	superHerUseCase domain.SuperHeroUseCase
	responder       *web.Responder
	encrypter       domain.Encrypter
}

func (controller *FetchSuperHeroController) Inject(responder *web.Responder, superHerUseCase domain.SuperHeroUseCase, encrypter domain.Encrypter) {
	controller.responder = responder
	controller.superHerUseCase = superHerUseCase
	controller.encrypter = encrypter
}

func NewFetchController(superHerUseCase domain.SuperHeroUseCase, encrypter domain.Encrypter) *FetchSuperHeroController {
	return &FetchSuperHeroController{
		superHerUseCase: superHerUseCase,
		encrypter:       encrypter,
	}
}

func (controller *FetchSuperHeroController) Wrapper(ctx context.Context, r *web.Request) web.Result {
	request, err := presentation_adapter.AdapterRequest[any](r)
	if err != nil {
		return controller.responder.ServerError(domain.BadRequest(err.Error()))
	}
	response := controller.Handle(ctx, request)
	if response.Error != nil {
		return controller.responder.ServerErrorWithCodeAndTemplate(response.Error, "error/withCode", response.StatusCode)
	}
	return controller.responder.Data(response.Data).Status(response.StatusCode)
}

func (controller *FetchSuperHeroController) isAdmin(ctx context.Context, request presentation.Request[any], data []domain.Superhero) ([]domain.SuperHeroWithEncryptIdentity, error) {
	var superHeroesParsed []domain.SuperHeroWithEncryptIdentity
	if request.Headers["X-Dee-See-Admin-Key"] == nil {
		for i := range data {
			identity, err := controller.superHerUseCase.EncryptIdentity(ctx, data[i].Identity)
			if err != nil {
				return nil, err
			}
			superHeroesParsed = append(superHeroesParsed, domain.SuperHeroWithEncryptIdentity{
				Name:        data[i].Name,
				Identity:    identity,
				Birthday:    data[i].Birthday,
				Superpowers: data[i].Superpowers,
			})
		}
	} else {
		superHeroesParsed = domain.ParseResponse(data)
	}
	return superHeroesParsed, nil
}

func (controller *FetchSuperHeroController) Handle(ctx context.Context, request presentation.Request[any]) presentation.Response[SuperHeroResponse] {
	if len(request.Query) > 0 {
		statusCode, heroes, err := controller.filter(request, ctx)
		if err != nil {
			return presentation.CreateResponse[SuperHeroResponse](uint(statusCode), nil, err)
		}
		superHeroesParsed, err := controller.isAdmin(ctx, request, heroes)
		return presentation.CreateResponse[SuperHeroResponse](uint(statusCode), superHeroesParsed, err)
	}
	data, err := controller.superHerUseCase.Fetch(ctx)
	if err != nil {
		return presentation.CreateResponse[SuperHeroResponse](http.StatusInternalServerError, nil, err)
	}
	superHeroesParsed, err := controller.isAdmin(ctx, request, data)
	if err != nil {
		return presentation.CreateResponse[SuperHeroResponse](http.StatusOK, nil, err)
	}
	return presentation.CreateResponse[SuperHeroResponse](http.StatusOK, superHeroesParsed, nil)
}

func (controller *FetchSuperHeroController) filter(request presentation.Request[any], ctx context.Context) (int, []domain.Superhero, error) {
	var data []domain.Superhero
	var err error
	for key, val := range request.Query {
		switch key {
		case presentation.SuperPowerFilter:
			data, err = controller.superHerUseCase.GetBySuperPower(ctx, val)
		default:
			return http.StatusBadRequest, nil, domain.BadRequest("invalid filter")
		}
	}
	return http.StatusOK, data, err
}

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
	adminAuth       domain.Authentication[map[string][]string, bool]
}

func (controller *FetchSuperHeroController) Inject(responder *web.Responder, superHerUseCase domain.SuperHeroUseCase, encrypter domain.Encrypter, adminAuth domain.Authentication[map[string][]string, bool]) {
	controller.responder = responder
	controller.superHerUseCase = superHerUseCase
	controller.encrypter = encrypter
	controller.adminAuth = adminAuth
}

func NewFetchController(superHerUseCase domain.SuperHeroUseCase, encrypter domain.Encrypter, adminAuth domain.Authentication[map[string][]string, bool]) *FetchSuperHeroController {
	return &FetchSuperHeroController{
		superHerUseCase: superHerUseCase,
		encrypter:       encrypter,
		adminAuth:       adminAuth,
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
	return controller.responder.Data(response).Status(response.StatusCode)
}

func (controller *FetchSuperHeroController) isAdmin(ctx context.Context, request presentation.Request[any], data []domain.Superhero) ([]domain.SuperHeroWithEncryptIdentity, error) {
	var superHeroesParsed []domain.SuperHeroWithEncryptIdentity
	auth, err := controller.adminAuth.Auth(request.Headers)
	if err != nil || !auth {
		for i := range data {
			identity, err := controller.superHerUseCase.EncryptIdentity(ctx, data[i].Identity)
			if err != nil {
				return nil, err
			}
			superHeroesParsed = append(superHeroesParsed, domain.SuperHeroWithEncryptIdentity{
				ID:          data[i].ID,
				Name:        data[i].Name,
				Identity:    identity,
				Birthday:    data[i].Birthday,
				Superpowers: data[i].Superpowers,
			})
		}
		return superHeroesParsed, nil
	}
	superHeroesParsed = domain.ParseResponse(data)
	return superHeroesParsed, nil
}

func (controller *FetchSuperHeroController) Handle(ctx context.Context, request presentation.Request[any]) presentation.Response[SuperHeroResponse] {
	var filter map[string][]string
	if len(request.Query) > 0 {
		filter = request.Query
	}
	data, err := controller.superHerUseCase.Fetch(ctx, filter)
	if err != nil {
		return presentation.CreateResponse[SuperHeroResponse](http.StatusInternalServerError, nil, err)
	}
	superHeroesParsed, err := controller.isAdmin(ctx, request, data)
	if err != nil {
		return presentation.CreateResponse[SuperHeroResponse](http.StatusOK, nil, err)
	}
	return presentation.CreateResponse[SuperHeroResponse](http.StatusOK, superHeroesParsed, nil)
}

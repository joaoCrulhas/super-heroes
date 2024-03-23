package presentation

import (
	"context"
	"net/http"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	presentation "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
	presentation_adapter "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation/controllers/adapter"
)

type CreateSuperHeroController struct {
	superHerUseCase domain.SuperHeroUseCase
	responder       *web.Responder
	auth            domain.Authentication[map[string][]string, bool]
	validators      []Validator
}

func (controller *CreateSuperHeroController) Inject(responder *web.Responder, superHerUseCase domain.SuperHeroUseCase, encrypter domain.Encrypter, auth domain.Authentication[map[string][]string, bool]) {
	controller.responder = responder
	controller.superHerUseCase = superHerUseCase
	controller.auth = auth
	controller.validators = append(controller.validators, ValidateSuperPowerInput)
}

func NewCreateController(superHerUseCase domain.SuperHeroUseCase, validators ...Validator) *CreateSuperHeroController {
	return &CreateSuperHeroController{
		superHerUseCase: superHerUseCase,
		validators:      validators,
	}
}

func (controller *CreateSuperHeroController) Wrapper(ctx context.Context, r *web.Request) web.Result {
	request, err := presentation_adapter.AdapterRequest[*domain.Superhero](r)
	if err != nil {
		return controller.responder.ServerError(err)
	}

	errValidators := controller.validateSuperPowers(request)
	if errValidators != nil {
		return controller.responder.ServerErrorWithCodeAndTemplate(errValidators, "error/withCode", http.StatusBadRequest)
	}
	auth, err := controller.auth.Auth(request.Headers)
	if err != nil || !auth {
		return controller.responder.ServerErrorWithCodeAndTemplate(err, "error/withCode", http.StatusUnauthorized)
	}
	response := controller.Handle(ctx, request)
	if response.Error != nil {
		return controller.responder.ServerErrorWithCodeAndTemplate(response.Error, "error/withCode", response.StatusCode)
	}
	return controller.responder.Data(response).Status(response.StatusCode)
}

func (controller *CreateSuperHeroController) validateSuperPowers(request presentation.Request[*domain.Superhero]) error {
	for _, value := range request.Body.Superpowers {
		for _, validator := range controller.validators {
			err := validator(value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (controller *CreateSuperHeroController) Handle(ctx context.Context, request presentation.Request[*domain.Superhero]) presentation.Response[domain.SuperHeroWithEncryptIdentity] {
	hero, err := controller.superHerUseCase.Create(ctx, request.Body)
	if err != nil {
		return presentation.CreateResponse[domain.SuperHeroWithEncryptIdentity](500, domain.SuperHeroWithEncryptIdentity{}, err)
	}
	return presentation.CreateResponse(http.StatusCreated, hero, nil)
}

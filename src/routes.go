package superhero

import (
	"flamingo.me/flamingo/v3/framework/web"
	controllers "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation/controllers"
)

type Routes struct {
	fetchAllController *controllers.FetchSuperHeroController
	createSuperHero    *controllers.CreateSuperHeroController
}

func (r *Routes) Inject(fetchAllController *controllers.FetchSuperHeroController, createSuperHero *controllers.CreateSuperHeroController) {
	r.fetchAllController = fetchAllController
	r.createSuperHero = createSuperHero
}

// Routes definition for the module
func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.HandlePost("superhero.create", r.createSuperHero.Wrapper)
	registry.HandleGet("superhero.all", r.fetchAllController.Wrapper)
	registry.Route("/super-heroes", "superhero.all")
	registry.Route("/super-heroes", "superhero.create")
}

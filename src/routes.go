package superhero

import (
	"flamingo.me/flamingo/v3/framework/web"
	controllers "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation/controllers"
)

type Routes struct {
	fetchAllController *controllers.FetchSuperHeroController
}

func (r *Routes) Inject(fetchAllController *controllers.FetchSuperHeroController) {
	r.fetchAllController = fetchAllController
}

// Routes definition for the module
func (r *Routes) Routes(registry *web.RouterRegistry) {
	registry.HandleGet("superhero.all", r.fetchAllController.Wrapper)
	registry.Route("/super-heroes", "superhero.all")
}

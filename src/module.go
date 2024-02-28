package superhero

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/infra/db"
	memory_db "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/db/memory"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/usecases"
)

// Module Basic struct
type Module struct{}

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	memoryDb := memory_db.NewSuperHeroMemoryRepository(nil)
	superHeroUseCases := usecases.NewSuperHeroUseCase(memoryDb)

	web.BindRoutes(injector, new(Routes))
	injector.Bind(new(db.Repository[domain.Superhero])).ToInstance(memoryDb)
	injector.Bind(new(domain.SuperHeroUseCase)).ToInstance(superHeroUseCases)
}

package superhero

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/infra/db"
	memory_db "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/db/memory"
	encrypter "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/deesee-chiffre"
	encrypter_validators "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/deesee-chiffre/validators"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/usecases"
)

// Module Basic struct
type Module struct{}

const (
	MinShiftValue = 96
	MaxShiftValue = 122
	key           = 5
)

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	memoryDb, err := memory_db.NewSuperHeroMemoryRepository(nil)
	if err != nil {
		panic(err)
	}
	adminAuth := usecases.NewAuthenticationAdmin()
	encrypter := encrypter.NewEncryptDeeSeeChiffreService(key, MinShiftValue, MaxShiftValue, encrypter_validators.ValidateEmptyInput, encrypter_validators.ValidateSpecialCharacters)
	superHeroUseCases := usecases.NewSuperHeroUseCase(memoryDb, encrypter)
	web.BindRoutes(injector, new(Routes))
	injector.Bind(new(domain.Authentication[map[string][]string, bool])).ToInstance(adminAuth)
	injector.Bind(new(db.Repository[domain.Superhero])).ToInstance(memoryDb)
	injector.Bind(new(domain.SuperHeroUseCase)).ToInstance(superHeroUseCases)
	injector.Bind(new(domain.Encrypter)).ToInstance(encrypter)
}

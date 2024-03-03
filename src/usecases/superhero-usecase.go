package usecases

import (
	"context"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/infra/db"
)

type superHeroUseCase struct {
	repository db.Repository[domain.Superhero]
	encrypter  domain.Encrypter
}

func NewSuperHeroUseCase(repository db.Repository[domain.Superhero], encrypter domain.Encrypter) *superHeroUseCase {
	return &superHeroUseCase{repository: repository, encrypter: encrypter}
}

func (su *superHeroUseCase) Inject(repository db.Repository[domain.Superhero], encrypter domain.Encrypter) {
	su.repository = repository
	su.encrypter = encrypter
}

func (su *superHeroUseCase) Create(ctx context.Context, superHero domain.Superhero) (domain.SuperHeroWithEncryptIdentity, error) {
	superHeroCreated, err := su.repository.Create(ctx, superHero)

	if err != nil {
		return domain.SuperHeroWithEncryptIdentity{}, err
	}
	return domain.ParseSuperHero(superHeroCreated), nil
}

func (su *superHeroUseCase) Fetch(ctx context.Context, filter map[string][]string) ([]domain.Superhero, error) {
	if filter == nil {
		return su.repository.Fetch(ctx)
	}
	return su.repository.FindByFilter(ctx, filter)
}

func (su *superHeroUseCase) GetBySuperPower(ctx context.Context, powers []string) ([]domain.Superhero, error) {
	return su.repository.FindByFilter(ctx, map[string][]string{"superpowers": powers})
}

func (su *superHeroUseCase) EncryptIdentity(ctx context.Context, identity domain.Identity) (string, error) {
	firstName, err := su.encrypter.Encrypt(identity.FirstName)
	if err != nil {
		return "", err
	}
	lastName, err := su.encrypter.Encrypt(identity.LastName)
	if err != nil {
		return "", err
	}

	return firstName + " " + lastName, nil
}

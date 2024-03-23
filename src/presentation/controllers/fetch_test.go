package presentation_test

import (
	"context"
	"testing"

	mocks "github.com/joaoCrulhas/omnevo-super-heroes/mocks/github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	shero_domain "github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
	controllers "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation/controllers"
	testutils "github.com/joaoCrulhas/omnevo-super-heroes/src/test-utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type FetchControllerTestSuite struct {
	suite.Suite
	sut           *controllers.FetchSuperHeroController
	mockedUseCase *mocks.MockSuperHeroUseCase
	mockEncrypter *mocks.MockEncrypter
	mockAdminAuth *mocks.MockAuthentication[map[string][]string, bool]
	ctx           context.Context
}

// this function executes before the test suite begins execution
func (suite *FetchControllerTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.mockAdminAuth = mocks.NewMockAuthentication[map[string][]string, bool](suite.T())
	suite.mockedUseCase = mocks.NewMockSuperHeroUseCase(suite.T())
	suite.mockEncrypter = mocks.NewMockEncrypter(suite.T())
	suite.sut = controllers.NewFetchController(suite.mockedUseCase, suite.mockEncrypter, suite.mockAdminAuth)
}

// This TestSuite is responsible for testing the FetchSuperHeroController
func (suite *FetchControllerTestSuite) TestShouldReturnAllHeroes() {
	expected := map[int]shero_domain.Superhero{}
	expected[1] = shero_domain.Superhero{
		Name: "superHero1",
		Identity: shero_domain.Identity{
			FirstName: "Snyder",
			LastName:  "Johnston",
		},
		Birthday:    "1990-04-14",
		Superpowers: []string{"flight", "strength", "invulnerability"},
	}
	expected[2] = shero_domain.Superhero{
		Name: "Super Hero 2",
		Identity: shero_domain.Identity{
			FirstName: "Snyder",
			LastName:  "Johnston",
		},
		Birthday:    "1973-04-18", // Batman's first appearance in comics
		Superpowers: []string{},
	}

	suite.mockedUseCase.EXPECT().Fetch(suite.ctx, mock.Anything).Return(expected, nil)
	suite.mockedUseCase.EXPECT().EncryptIdentity(suite.ctx, mock.Anything).Return("mock", nil)
	suite.mockAdminAuth.EXPECT().Auth(mock.Anything).Return(false, shero_domain.Unauthorized("unauthorized"))
	request := presentation.Request[any]{}
	actual := suite.sut.Handle(suite.ctx, request)
	suite.Assertions.Equal(shero_domain.SuperHeroWithEncryptIdentity{
		Name:        "superHero1",
		Identity:    "mock",
		Birthday:    "1990-04-14",
		Superpowers: []string{"flight", "strength", "invulnerability"},
	}, actual.Data[0])

	suite.Assertions.Equal(shero_domain.SuperHeroWithEncryptIdentity{
		Name:        "Super Hero 2",
		Identity:    "mock",
		Birthday:    "1973-04-18", // Batman's first appearance in comics
		Superpowers: []string{},
	}, actual.Data[1])
	suite.Assertions.Equal(uint(200), actual.StatusCode)
}

// This TestSuite is responsible for testing the opportunity to filter the heroes by different super powers.
func (suite *FetchControllerTestSuite) TestShouldReturnHeroesWithDifferentSuperPowers() {
	suite.mockedUseCase.EXPECT().Fetch(suite.ctx, mock.Anything).Return(testutils.GetSuperHeroes(), nil)
	suite.mockedUseCase.EXPECT().EncryptIdentity(suite.ctx, mock.Anything).Return("mock", nil)
	request := presentation.Request[any]{
		Query: map[string][]string{
			"superpowers": {"strength", "healing"},
		},
	}
	actual := suite.sut.Handle(suite.ctx, request)
	suite.Assertions.Equal(uint(200), actual.StatusCode)
}

func TestFetchControllerTestSuite(t *testing.T) {
	suite.Run(t, new(FetchControllerTestSuite))
}

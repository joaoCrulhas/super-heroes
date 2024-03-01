package presentation_test

import (
	"context"
	"errors"
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
	ctx           context.Context
}

// this function executes before the test suite begins execution
func (suite *FetchControllerTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.mockedUseCase = mocks.NewMockSuperHeroUseCase(suite.T())
	suite.mockEncrypter = mocks.NewMockEncrypter(suite.T())
	suite.sut = controllers.NewFetchController(suite.mockedUseCase, suite.mockEncrypter)
}

func (suite *FetchControllerTestSuite) TestShouldReturnAllHeroes() {
	expected := []shero_domain.Superhero{
		{
			Name: "superHero1",
			Identity: shero_domain.Identity{
				FirstName: "Snyder",
				LastName:  "Johnston",
			},
			Birthday:    "1990-04-14",
			Superpowers: []string{"flight", "strength", "invulnerability"},
		},
		{
			Name: "Super Hero 2",
			Identity: shero_domain.Identity{
				FirstName: "Snyder",
				LastName:  "Johnston",
			},
			Birthday:    "1973-04-18", // Batman's first appearance in comics
			Superpowers: []string{},
		},
	}
	suite.mockedUseCase.EXPECT().Fetch(suite.ctx).Return(expected, nil).Once()
	suite.mockedUseCase.EXPECT().EncryptIdentity(suite.ctx, mock.Anything).Return("mock", nil)
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

func (suite *FetchControllerTestSuite) TestShouldReturnStatusCode500IfError() {
	suite.mockedUseCase.EXPECT().Fetch(suite.ctx).Return(nil, errors.New("error")).Once()
	request := presentation.Request[any]{}
	actual := suite.sut.Handle(suite.ctx, request)
	suite.Assertions.Equal(uint(500), actual.StatusCode)
}

func (suite *FetchControllerTestSuite) TestShouldReturnHeroesFilteredBySuperPower() {
	suite.mockedUseCase.EXPECT().GetBySuperPower(suite.ctx, mock.Anything).Return(testutils.GetSuperHeroes(), nil).Once()
	request := presentation.Request[any]{
		Query: map[string][]string{
			"superpowers": {"strength"},
		},
	}
	actual := suite.sut.Handle(suite.ctx, request)
	suite.Assertions.Equal(uint(200), actual.StatusCode)
}

func (suite *FetchControllerTestSuite) TestShouldReturnHeroesWithDifferentSuperPowers() {
	suite.mockedUseCase.EXPECT().GetBySuperPower(suite.ctx, mock.Anything).Return(testutils.GetSuperHeroes(), nil).Once()
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

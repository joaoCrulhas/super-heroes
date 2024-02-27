package presentation_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	mocks "github.com/joaoCrulhas/omnevo-super-heroes/mocks/github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	"github.com/joaoCrulhas/omnevo-super-heroes/src/presentation"
	controllers "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation/controllers"
	testutils "github.com/joaoCrulhas/omnevo-super-heroes/src/test-utils"
	"github.com/stretchr/testify/suite"
)

type FetchControllerTestSuite struct {
	suite.Suite
	sut           *controllers.FetchSuperHeroController
	mockedUseCase *mocks.MockSuperHeroUseCase
	ctx           context.Context
}

// this function executes before the test suite begins execution
func (suite *FetchControllerTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	suite.mockedUseCase = mocks.NewMockSuperHeroUseCase(suite.T())
	suite.sut = controllers.NewFetchController(suite.mockedUseCase)
}

// this function executes after all tests executed
func (suite *FetchControllerTestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

// this function executes before each test case
func (suite *FetchControllerTestSuite) SetupTest() {
	fmt.Println("-- From SetupTest")
}

// this function executes after each test case
func (suite *FetchControllerTestSuite) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func (suite *FetchControllerTestSuite) TestShouldReturnAllHeroes() {
	expected := testutils.GetSuperHeroes()
	suite.mockedUseCase.EXPECT().Fetch(suite.ctx).Return(expected, nil).Once()
	request := presentation.Request[any]{}
	actual := suite.sut.Handle(suite.ctx, request)
	suite.Equal(expected, actual.Data)
	suite.Equal(200, actual.StatusCode)
}

func (suite *FetchControllerTestSuite) TestShouldReturnStatusCode500IfError() {
	suite.mockedUseCase.EXPECT().Fetch(suite.ctx).Return(nil, errors.New("error")).Once()
	request := presentation.Request[any]{}
	actual := suite.sut.Handle(suite.ctx, request)
	suite.Assertions.Equal(500, actual.StatusCode)
}

func TestFetchControllerTestSuite(t *testing.T) {
	suite.Run(t, new(FetchControllerTestSuite))
}

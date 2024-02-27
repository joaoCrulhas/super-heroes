package presentation_test

import (
	"fmt"
	"testing"

	controllers "github.com/joaoCrulhas/omnevo-super-heroes/src/presentation/controllers"
	"github.com/stretchr/testify/suite"
)

type FetchControllerTestSuite struct {
	suite.Suite
	sut *controllers.FetchSuperHeroController
}

// this function executes before the test suite begins execution
func (suite *FetchControllerTestSuite) SetupSuite() {
	fmt.Println(">>> From SetupSuite")
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

func (suite *FetchControllerTestSuite) TestShouldCall() {
	fmt.Println("From TestExample")
}

func TestFetchControllerTestSuite(t *testing.T) {
	suite.Run(t, new(FetchControllerTestSuite))
}

package encrypter_test

import (
	"fmt"
	"testing"

	encrypter "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/service"
	"github.com/stretchr/testify/suite"
)

type EncrypterServiceSuite struct {
	suite.Suite
	sut encrypter.Service
}

// this function executes before the test suite begins execution
func (suite *EncrypterServiceSuite) SetupSuite() {
	// set StartingNumber to one
	fmt.Println(">>> From SetupSuite")
	suite.sut = encrypter.NewEncryptService()
}

// this function executes after all tests executed
func (suite *EncrypterServiceSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

// this function executes before each test case
func (suite *EncrypterServiceSuite) SetupTest() {
	// reset StartingNumber to one
	fmt.Println("-- From SetupTest")
}

// this function executes after each test case
func (suite *EncrypterServiceSuite) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func (suite *EncrypterServiceSuite) TestShouldReturnErrorIfInputIsEmpty() {
	_, err := suite.sut.Encrypt("")
	suite.Assertions.Error(err)
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(EncrypterServiceSuite))
}

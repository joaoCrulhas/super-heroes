package encrypter_test

import (
	"fmt"
	"testing"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/infra/dictionary"
	encrypter "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/key-strategy"
	validators "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/key-strategy/validators"
	"github.com/stretchr/testify/suite"
)

type EncrypterServiceSuite struct {
	suite.Suite
	sut *encrypter.Service
}

const key = 5

// this function executes before the test suite begins execution
func (suite *EncrypterServiceSuite) SetupSuite() {
	fmt.Println(">>> From SetupSuite")
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	dictionary := dictionary.NewDictionaryIsoAlphabetic(alphabet, dictionary.Compute(alphabet))
	suite.sut = encrypter.NewEncryptService(key, dictionary, validators.ValidateEmptyInput, validators.ValidateSpecialCharacters)
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
func (suite *EncrypterServiceSuite) TestShouldReturnErrorIfHasInvalidCharacters() {
	_, err := suite.sut.Encrypt("ab@#$!c123")
	suite.Assertions.Error(err)
}

func (suite *EncrypterServiceSuite) TestShouldReturnAnEncryptStringIfTheInputIsCorrect() {
	expected := "otft"
	got, _ := suite.sut.Encrypt("joao")
	suite.Assertions.Equal(got, expected)
}

func (suite *EncrypterServiceSuite) TestShouldReturnAnEncryptedStringIfTheInputExceedsTheAlphabet() {
	expected := "bcde"
	got, _ := suite.sut.Encrypt("wxyz")
	suite.Assertions.Equal(got, expected)
}

func TestEncrypterAlphabeticIso(t *testing.T) {
	suite.Run(t, new(EncrypterServiceSuite))
}

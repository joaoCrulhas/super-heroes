package encrypter_test

import (
	"testing"

	encrypter "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/deesee-chiffre"
	validators "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/encrypter/deesee-chiffre/validators"
	"github.com/stretchr/testify/suite"
)

type EncrypterServiceSuite struct {
	suite.Suite
	sut *encrypter.EncryptDeeSeeChiffreService
}

const (
	MinShiftValue = 96
	MaxShiftValue = 122
	key           = 5
)

// this function executes before the test suite begins execution
func (suite *EncrypterServiceSuite) SetupSuite() {
	suite.sut = encrypter.NewEncryptDeeSeeChiffreService(key, MinShiftValue, MaxShiftValue, validators.ValidateEmptyInput, validators.ValidateSpecialCharacters)
}

// this function executes after all tests executed
func (suite *EncrypterServiceSuite) TearDownSuite() {
}

// this function executes before each test case
func (suite *EncrypterServiceSuite) SetupTest() {
}

// this function executes after each test case
func (suite *EncrypterServiceSuite) TearDownTest() {
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
	expected := "hqfwp"
	actual, _ := suite.sut.Encrypt("clark")
	suite.Assertions.Equal(expected, actual)
}

func TestShouldReturnACorrectValueWithKeyEqualsThree(t *testing.T) {
	sut := encrypter.NewEncryptDeeSeeChiffreService(3, MinShiftValue, MaxShiftValue, validators.ValidateEmptyInput, validators.ValidateSpecialCharacters)
	actual1, _ := sut.Encrypt("cherry")
	if actual1 != "fkhuub" {
		t.Fatalf(`want %q, got %v error`, "b", actual1)
	}
	actual2, _ := sut.Encrypt("blossom")
	if actual2 != "eorvvrp" {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, "eoorvrf", actual2)
	}
}

func TestShouldReturnACorrectValueWithKeyEqualsTwentySix(t *testing.T) {
	sut := encrypter.NewEncryptDeeSeeChiffreService(26, MinShiftValue, MaxShiftValue, validators.ValidateEmptyInput, validators.ValidateSpecialCharacters)
	actual1, _ := sut.Encrypt("zyxwvutsrqponmlkjihgfedcba")
	if actual1 != "zyxwvutsrqponmlkjihgfedcba" {
		t.Fatalf(`want %q, got %v error`, "zyxwvutsrqponmlkjihgfedcba", actual1)
	}
}

func TestEncrypterAlphabeticIso(t *testing.T) {
	suite.Run(t, new(EncrypterServiceSuite))
}

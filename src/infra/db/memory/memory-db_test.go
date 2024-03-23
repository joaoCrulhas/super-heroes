package db_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	db "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/db/memory"
	testutils "github.com/joaoCrulhas/omnevo-super-heroes/src/test-utils"
	"github.com/stretchr/testify/suite"
)

type MemoryDbTestSuite struct {
	suite.Suite
	sut *db.SuperHeroMemoryRepository
	ctx context.Context
}

// this function executes before the test suite begins execution
func (suite *MemoryDbTestSuite) SetupSuite() {
	fmt.Println(">>> From SetupSuite")
	sut, err := db.NewSuperHeroMemoryRepository(testutils.GetSuperHeroes())
	if err != nil {
		suite.T().Errorf("Error creating SuperHeroMemoryRepository: %v", err)
	}
	suite.sut = sut
	suite.ctx = context.Background()
}

// this function executes after all tests executed
func (suite *MemoryDbTestSuite) TearDownSuite() {
	fmt.Println(">>> From TearDownSuite")
}

// this function executes before each test case
func (suite *MemoryDbTestSuite) SetupTest() {
	fmt.Println("-- From SetupTest")
}

// this function executes after each test case
func (suite *MemoryDbTestSuite) TearDownTest() {
	fmt.Println("-- From TearDownTest")
}

func (suite *MemoryDbTestSuite) TestShouldReturnAllHeroes() {
	got, _ := suite.sut.Fetch(suite.ctx)
	suite.Equal(testutils.GetSuperHeroes(), got)
}

func (suite *MemoryDbTestSuite) TestUsingFindByFilter() {
	filter := map[string][]string{"superpowers": {"strength"}}
	expected := testutils.GetSuperHeroes()[1]
	got, _ := suite.sut.FindByFilter(suite.ctx, filter)
	suite.Equal(expected, got[1])
}

func (suite *MemoryDbTestSuite) TestShouldMatchTwoSuperPowers() {
	filter := map[string][]string{"superpowers": {"strength", "healing"}}
	exp := domain.SuperHerosData{}
	exp[1] = &domain.Superhero{
		ID:   1,
		Name: "superHero1",
		Identity: domain.Identity{
			FirstName: "Snyder",
			LastName:  "Johnston",
		},
		Birthday:    "1990-04-14",
		Superpowers: []string{"flight", "strength", "invulnerability"},
	}
	exp[3] = &domain.Superhero{
		ID:   3,
		Name: "superHero3",
		Identity: domain.Identity{
			FirstName: "Test3",
			LastName:  "TestLastName3",
		},
		Birthday:    "1990-04-14", // Batman's first appearance in comics
		Superpowers: []string{"healing"},
	}
	got, _ := suite.sut.FindByFilter(suite.ctx, filter)
	suite.Equal(exp, got)
}

func (suite *MemoryDbTestSuite) TestShouldReturnASuperHeroIfMatches() {
	filter := map[string][]string{"superpowers": {"strength", "healing"}}

	got, _ := suite.sut.FindByFilter(suite.ctx, filter)
	suite.Equal(&domain.Superhero{
		ID:   1,
		Name: "superHero1",
		Identity: domain.Identity{
			FirstName: "Snyder",
			LastName:  "Johnston",
		},
		Birthday:    "1990-04-14",
		Superpowers: []string{"flight", "strength", "invulnerability"},
	}, got[1])
	suite.Equal(&domain.Superhero{
		ID:          3,
		Name:        "superHero3",
		Birthday:    "1990-04-14",
		Superpowers: []string{"healing"},
		Identity: domain.Identity{
			FirstName: "Test3",
			LastName:  "TestLastName3",
		},
	}, got[3])
}

func (suite *MemoryDbTestSuite) TestShouldReturnAnEmptyArrayIfNoSuperHeroWithSuperPower() {
	filter := map[string][]string{"superpowers": {"invisibility"}}
	got, _ := suite.sut.FindByFilter(suite.ctx, filter)
	suite.Equal(0, len(got))
}

func (suite *MemoryDbTestSuite) TestShouldNotReturnRepetitiveSuperHero() {
	filter := map[string][]string{"superpowers": {"invisibility", "healing"}}
	mockMap := domain.SuperHerosData{}
	mockMap[1] = &domain.Superhero{
		Name:        "superHero1",
		ID:          1,
		Birthday:    "1990-04-14",
		Superpowers: []string{"healing", "invisibility"},
		Identity: domain.Identity{
			FirstName: "Test",
			LastName:  "Two",
		},
	}
	sut, _ := db.NewSuperHeroMemoryRepository(mockMap)
	got, _ := sut.FindByFilter(suite.ctx, filter)
	suite.Equal(&domain.Superhero{
		ID:   1,
		Name: "superHero1",
		Identity: domain.Identity{
			FirstName: "Test",
			LastName:  "Two",
		},
		Birthday:    "1990-04-14",
		Superpowers: []string{"healing", "invisibility"},
	}, got[1])

}

func TestMemoryDbTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryDbTestSuite))
}

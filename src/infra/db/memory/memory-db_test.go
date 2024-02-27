package db_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
	db "github.com/joaoCrulhas/omnevo-super-heroes/src/infra/db/memory"
	"github.com/stretchr/testify/suite"
)

type MemoryDbTestSuite struct {
	suite.Suite
	sut db.SuperHeroMemoryRepository
	ctx context.Context
}

// this function executes before the test suite begins execution
func (suite *MemoryDbTestSuite) SetupSuite() {
	fmt.Println(">>> From SetupSuite")
	suite.sut = db.NewSuperHeroMemoryRepository(getSuperHeroes())
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
	fmt.Println("From TestExample")
	expected := getSuperHeroes()
	got, _ := suite.sut.Fetch(suite.ctx)
	suite.Equal(expected, got)
}

func (suite *MemoryDbTestSuite) TestUsingFindByFilter() {
	filter := map[string]any{"superpowers": "flight"}
	fmt.Println("From TestExample")
	expected := getSuperHeroes()[0:1]
	got, _ := suite.sut.FindByFilter(suite.ctx, filter)
	suite.Equal(expected, got)
}

func (suite *MemoryDbTestSuite) TestShouldReturnAnEmptyArrayIfNoSuperHeroWithSuperPower() {
	filter := map[string]any{"superpowers": "random_superpower"}
	fmt.Println("From TestExample")
	got, _ := suite.sut.FindByFilter(suite.ctx, filter)
	suite.Equal([]domain.Superhero(nil), got)
}

func TestMemoryDbTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryDbTestSuite))
}

func getSuperHeroes() []domain.Superhero {
	heroes := []domain.Superhero{ // Change from array to slice
		{
			Name: "superHero1",
			Identity: domain.Identity{
				FirstName: "Snyder",
				LastName:  "Johnston",
			},
			Birthday:    "1990-04-14",
			Superpowers: []string{"flight", "strength", "invulnerability"},
		},
		{
			Name: "Super Hero 2",
			Identity: domain.Identity{
				FirstName: "Petra",
				LastName:  "Sharpe",
			},
			Birthday:    "1973-04-18", // Batman's first appearance in comics
			Superpowers: []string{},
		},
	}
	return heroes
}

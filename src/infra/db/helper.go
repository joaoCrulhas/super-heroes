package db

import (
	"encoding/json"
	"os"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
)

func ReadSuperHeroFile(fileName string) ([]domain.Superhero, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var superHeroes []domain.Superhero
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&superHeroes)
	if err != nil {
		return nil, err
	}
	return superHeroes, nil
}

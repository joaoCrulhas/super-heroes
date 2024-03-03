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

	for i := range superHeroes {
		superHeroes[i].ID = uint64(i + 1)
	}
	return superHeroes, nil
}

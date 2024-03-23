package db

import (
	"encoding/json"
	"os"

	"github.com/joaoCrulhas/omnevo-super-heroes/src/domain"
)

func ReadSuperHeroFile(fileName string) (map[int]domain.Superhero, error) {
	superHeroes := map[int]domain.Superhero{}
	var parsed []domain.Superhero
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&parsed)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(parsed); i++ {
		parsed[i].ID = i + 1
		superHeroes[parsed[i].ID] = parsed[i]
	}
	return superHeroes, nil
}

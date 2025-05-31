package repository

import (
	"encoding/json"
	"os"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type Monsters struct {
	Monsters []entity.Monster `json:"monsters"`
}

type InMemoryMonsterRepository struct {
	monsters []entity.Monster
}

func NewInMemoryMonsterRepository() *InMemoryMonsterRepository {
	return &InMemoryMonsterRepository{
		monsters: []entity.Monster{},
	}
}

func (repo *InMemoryMonsterRepository) Save(monster entity.Monster) error {
	repo.monsters = append(repo.monsters, monster)
	return nil
}

func (repo *InMemoryMonsterRepository) Find(name string) (*entity.Monster, error) {
	for _, monster := range repo.monsters {
		if monster.Name == name {
			return &monster, nil
		}
	}
	return nil, nil
}

func (repo *InMemoryMonsterRepository) SaveJSONFile(fileName string) error {
	var monstersData Monsters
	monstersData.Monsters = repo.monsters

	monsterBytes, err := json.Marshal(monstersData)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, monsterBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (repo *InMemoryMonsterRepository) LoadJSONFile(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	var monstersData Monsters
	err = json.Unmarshal(file, &monstersData)
	if err != nil {
		return err
	}

	repo.monsters = monstersData.Monsters
	return nil
}

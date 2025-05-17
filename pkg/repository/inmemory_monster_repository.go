package repository

import (
	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

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

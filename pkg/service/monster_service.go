package service

import (
	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
)

type MonsterCheckService struct {
	monsterRepository repository.MonsterGetter
}

func NewMonsterService(monsterRepository repository.MonsterGetter) *MonsterCheckService {
	return &MonsterCheckService{
		monsterRepository: monsterRepository,
	}
}

func (s *MonsterCheckService) Exists(monster entity.Monster) (bool, error) {
	gotMonster, err := s.monsterRepository.Find(monster.Name)
	if err != nil {
		return false, err
	}
	if gotMonster == nil {
		return false, nil
	}

	return true, nil
}

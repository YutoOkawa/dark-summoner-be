package service

import (
	"errors"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
)

type MonsterGetInfoService struct {
	monsterRepository repository.MonsterGetter
}

func NewMonsterGetInfoService(monsterRepository repository.MonsterGetter) MonsterGetInfoService {
	return MonsterGetInfoService{
		monsterRepository: monsterRepository,
	}
}

func (s *MonsterGetInfoService) GetInfo(name string) (*entity.Monster, error) {
	monster, err := s.monsterRepository.Find(name)
	if err != nil {
		return nil, err
	}

	if monster == nil {
		return nil, errors.New("monster not found")
	}

	return monster, nil
}

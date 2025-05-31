package service

import (
	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
)

type MonsterListService struct {
	monsterReposiotry repository.MonsterLister
}

func NewMonsterListService(monsterReposiotry repository.MonsterLister) MonsterListService {
	return MonsterListService{
		monsterReposiotry: monsterReposiotry,
	}
}

func (s MonsterListService) List() ([]entity.Monster, error) {
	monsters, err := s.monsterReposiotry.List()
	if err != nil {
		return nil, err
	}
	return monsters, nil
}

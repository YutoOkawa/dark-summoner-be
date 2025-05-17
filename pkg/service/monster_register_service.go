package service

import (
	"errors"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
)

type MonsterRegisterService struct {
	monsterRepository repository.MonsterSaver
	monsterService    *MonsterCheckService
}

func NewMonsterRegisterService(monsterRepository repository.MonsterSaver, monsterService *MonsterCheckService) *MonsterRegisterService {
	return &MonsterRegisterService{
		monsterRepository: monsterRepository,
		monsterService:    monsterService,
	}
}

func (s *MonsterRegisterService) Register(monster entity.Monster) error {
	exists, err := s.monsterService.Exists(monster)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("monster already exists")
	}

	err = s.monsterRepository.Save(monster)
	if err != nil {
		return err
	}
	return nil
}

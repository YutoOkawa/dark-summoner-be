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

func NewMonsterRegisterService(monsterRepository repository.MonsterSaver, monsterService *MonsterCheckService) MonsterRegisterService {
	return MonsterRegisterService{
		monsterRepository: monsterRepository,
		monsterService:    monsterService,
	}
}

func (s *MonsterRegisterService) Register(command entity.MonsterRegisterCommand) error {
	atkMonsterParameter, err := entity.NewMonsterParameter(command.ATK)
	if err != nil {
		return err
	}

	defMonsterParameter, err := entity.NewMonsterParameter(command.DEF)
	if err != nil {
		return err
	}

	hpMonsterParameter, err := entity.NewMonsterParameter(command.HP)
	if err != nil {
		return err
	}

	monster, err := entity.NewMonster(
		command.Name,
		*atkMonsterParameter,
		*defMonsterParameter,
		*hpMonsterParameter,
		command.Rarity,
		command.RarityID,
		command.Skill,
		command.Attribute,
	)
	if err != nil {
		return err
	}

	exists, err := s.monsterService.Exists(*monster)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("monster already exists")
	}

	err = s.monsterRepository.Save(*monster)
	if err != nil {
		return err
	}
	return nil
}

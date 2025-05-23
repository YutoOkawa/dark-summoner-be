package service

import (
	"errors"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
)

type SummonerAddMonsterService struct {
	summonerRepository repository.SummonerGetterSaver
}

func NewSummonerAddMonsterService(summonerRepository repository.SummonerGetterSaver) SummonerAddMonsterService {
	return SummonerAddMonsterService{
		summonerRepository: summonerRepository,
	}
}

func (s *SummonerAddMonsterService) AddMonster(command entity.SummonerAddMonsterCommand) error {
	summoner, err := s.summonerRepository.Find(command.PlayerID)
	if err != nil {
		return err
	}

	if summoner == nil {
		return errors.New("summoner not found")
	}

	summoner.Monsters = append(summoner.Monsters, command.Monster)
	err = s.summonerRepository.Save(*summoner)
	if err != nil {
		return err
	}

	return nil
}

package service

import (
	"errors"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
)

type SummonerRegisterService struct {
	summonerRepository repository.SummonerSaver
	summonerService    *SummonerService
}

func NewSummonerRegisterService(summonerRepository repository.SummonerSaver, summonerService *SummonerService) SummonerRegisterService {
	return SummonerRegisterService{
		summonerRepository: summonerRepository,
		summonerService:    summonerService,
	}
}

func (s *SummonerRegisterService) Register(command entity.SummonerRegisterCommand) error {
	exists, err := s.summonerService.Exists(command.PlayerID)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("summoner already exists")
	}

	summoner, err := entity.NewSummoner(command.PlayerID)
	if err != nil {
		return err
	}

	summoner.Monsters = command.Monsters

	err = s.summonerRepository.Save(*summoner)
	if err != nil {
		return err
	}

	return nil
}

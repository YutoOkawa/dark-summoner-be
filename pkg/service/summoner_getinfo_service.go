package service

import (
	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
)

type SummonerGetInfoService struct {
	summonerRepository repository.SummonerGetter
}

func NewSummonerGetInfoService(summonerRepository repository.SummonerGetter) SummonerGetInfoService {
	return SummonerGetInfoService{
		summonerRepository: summonerRepository,
	}
}

func (s *SummonerGetInfoService) GetInfo(playerID string) (*entity.Summoner, error) {
	summoner, err := s.summonerRepository.Find(playerID)
	if err != nil {
		return nil, err
	}

	if summoner == nil {
		return nil, nil
	}

	return summoner, nil
}

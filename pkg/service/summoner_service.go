package service

import "github.com/YutoOkawa/dark-summoner-be/pkg/repository"

type SummonerService struct {
	repository repository.SummonerGetter
}

func NewSummonerService(summonerRepository repository.SummonerGetter) *SummonerService {
	return &SummonerService{
		repository: summonerRepository,
	}
}

func (s *SummonerService) Exists(playerID string) (bool, error) {
	summoner, err := s.repository.Find(playerID)
	if err != nil {
		return false, err
	}
	if summoner == nil {
		return false, nil
	}
	return true, nil
}

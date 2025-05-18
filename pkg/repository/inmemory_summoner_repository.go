package repository

import (
	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type InMemorySummonerRepository struct {
	summoners []entity.Summoner
}

func NewInMemorySummonerRepository() *InMemorySummonerRepository {
	return &InMemorySummonerRepository{
		summoners: []entity.Summoner{},
	}
}

func (repo *InMemorySummonerRepository) Save(summoner entity.Summoner) error {
	for i, s := range repo.summoners {
		if s.PlayerID == summoner.PlayerID {
			repo.summoners[i] = summoner
			return nil
		}
	}
	repo.summoners = append(repo.summoners, summoner)
	return nil
}

func (repo *InMemorySummonerRepository) Find(playerID string) (*entity.Summoner, error) {
	for _, summoner := range repo.summoners {
		if summoner.PlayerID == playerID {
			return &summoner, nil
		}
	}
	return nil, nil
}

package repository

import (
	"encoding/json"
	"os"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type Summoners struct {
	Summoners []entity.Summoner `json:"summoners"`
}

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
			// Update existing summoner
			repo.summoners[i] = summoner
			return nil
		}
	}
	// Add new summoner
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

func (repo *InMemorySummonerRepository) SaveJSONFile(fileName string) error {
	var summonersData Summoners
	summonersData.Summoners = repo.summoners

	summonerBytes, err := json.Marshal(summonersData)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, summonerBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (repo *InMemorySummonerRepository) LoadJSONFile(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	var summonersData Summoners
	err = json.Unmarshal(file, &summonersData)
	if err != nil {
		return err
	}
	repo.summoners = summonersData.Summoners
	return nil
}

package repository

import (
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

func TestInMemorySummonerRepositorySave(t *testing.T) {
	tests := []struct {
		name       string
		repository *InMemorySummonerRepository

		summoner entity.Summoner

		expectedError bool
	}{
		{
			name:       "SaveNewSummonerSuccessfully",
			repository: NewInMemorySummonerRepository(),

			summoner: entity.Summoner{
				PlayerID: "test_player_id",
				Monsters: []entity.Monster{},
			},

			expectedError: false,
		},
		{
			name: "SaveExistingSummoner",
			repository: func() *InMemorySummonerRepository {
				repo := NewInMemorySummonerRepository()
				repo.summoners = append(repo.summoners, entity.Summoner{
					PlayerID: "test_player_id",
					Monsters: []entity.Monster{},
				})
				return repo
			}(),

			summoner: entity.Summoner{
				PlayerID: "test_player_id",
				Monsters: []entity.Monster{
					{
						Name: "test_monster",
					},
				},
			},

			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemorySummonerRepository()
			err := repo.Save(tt.summoner)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if len(repo.summoners) != 1 {
				t.Errorf("expected 1 summoner, got: %d", len(repo.summoners))
			}
		})
	}
}

func TestInMemorySummonerRepositoryFind(t *testing.T) {
	test := []struct {
		name       string
		repository *InMemorySummonerRepository

		playerID string

		expectedSummoner *entity.Summoner
		expectedError    bool
	}{
		{
			name: "FindExistingSummonerSuccessfully",
			repository: func() *InMemorySummonerRepository {
				repo := NewInMemorySummonerRepository()
				repo.summoners = append(repo.summoners, entity.Summoner{
					PlayerID: "test_player_id",
					Monsters: []entity.Monster{},
				})
				return repo
			}(),

			playerID: "test_player_id",
			expectedSummoner: &entity.Summoner{
				PlayerID: "test_player_id",
				Monsters: []entity.Monster{},
			},
			expectedError: false,
		},
		{
			name:       "FindNonExistingSummoner",
			repository: NewInMemorySummonerRepository(),

			playerID: "non_existing_player_id",

			expectedSummoner: nil,
			expectedError:    false,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			summoner, err := tt.repository.Find(tt.playerID)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if summoner != nil && !summoner.Equal(tt.expectedSummoner) {
				t.Errorf("expected summoner: %v, got: %v", tt.expectedSummoner, summoner)
			}
		})
	}
}

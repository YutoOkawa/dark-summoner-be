package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

func TestSummonerGetInfoServiceGetInfo(t *testing.T) {
	tests := []struct {
		name string

		mockSummoner *entity.Summoner
		findError    error

		playerID string

		expectedSummoner *entity.Summoner
		expectedError    bool
	}{
		{
			name: "ShouldGetInfoSuccessfully",

			mockSummoner: &entity.Summoner{
				PlayerID: "test_player",
				Monsters: []entity.Monster{},
			},
			findError: nil,

			playerID: "test_player",

			expectedSummoner: &entity.Summoner{
				PlayerID: "test_player",
				Monsters: []entity.Monster{},
			},
			expectedError: false,
		},
		{
			name: "ShouldReturnErrorWhenSummonerDoesNotExist",

			mockSummoner: nil,
			findError:    nil,

			playerID: "non_existent_player",

			expectedSummoner: nil,
			expectedError:    false,
		},
		{
			name: "ShouldReturnErrorWhenRepositoryReturnsError",

			mockSummoner: &entity.Summoner{
				PlayerID: "test_player",
				Monsters: []entity.Monster{},
			},
			findError: errors.New("repository error"),

			playerID: "test_player",

			expectedSummoner: nil,
			expectedError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mockSummonerGetter{
				mockSummoner: tt.mockSummoner,
				findError:    tt.findError,
			}
			service := NewSummonerGetInfoService(&repo)

			summoner, err := service.GetInfo(tt.playerID)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if summoner != nil {
				if summoner.Equal(tt.expectedSummoner) == false {
					t.Errorf("expected summoner: %v, got: %v", tt.expectedSummoner, summoner)
				}
			}
		})
	}
}

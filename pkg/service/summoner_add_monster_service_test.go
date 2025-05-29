package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type mockSummonerGetterSaver struct {
	mockSummoner *entity.Summoner
	findError    error
	saveError    error
}

func (m *mockSummonerGetterSaver) Find(playerID string) (*entity.Summoner, error) {
	return m.mockSummoner, m.findError
}

func (m *mockSummonerGetterSaver) Save(summoner entity.Summoner) error {
	return m.saveError
}

func TestSummonerAddMonsterService(t *testing.T) {
	addMonsterCommand := entity.SummonerAddMonsterCommand{
		PlayerID: "test_player_id",
		Monster: entity.Monster{
			Name: "test_monster",
		},
	}

	tests := []struct {
		name string

		command entity.SummonerAddMonsterCommand

		mockSummoner *entity.Summoner
		findError    error
		saveError    error

		expectedError bool
	}{
		{
			name: "ShouldAddMonsterSuccessfully",

			command: addMonsterCommand,

			mockSummoner: &entity.Summoner{
				PlayerID: "test_player_id",
				Monsters: []entity.Monster{},
			},
			findError: nil,
			saveError: nil,

			expectedError: false,
		},
		{
			name: "ShouldReturnErrorWhenSummonerNotFound",

			command: addMonsterCommand,

			mockSummoner: nil,
			findError:    nil,
			saveError:    nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenFindError",

			command: addMonsterCommand,

			mockSummoner: nil,
			findError:    errors.New("find error"),
			saveError:    nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenSaveError",

			command: addMonsterCommand,

			mockSummoner: nil,
			findError:    nil,
			saveError:    errors.New("save error"),

			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockSummonerGetterSaver := &mockSummonerGetterSaver{
				mockSummoner: tt.mockSummoner,
				findError:    tt.findError,
				saveError:    tt.saveError,
			}
			summonerAddMonsterService := NewSummonerAddMonsterService(mockSummonerGetterSaver)

			err := summonerAddMonsterService.AddMonster(tt.command)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}

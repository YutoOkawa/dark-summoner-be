package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type mockSummonerSaver struct {
	saveError error
}

func (m *mockSummonerSaver) Save(summoner entity.Summoner) error {
	return m.saveError
}

func TestSummonerRegisterService(t *testing.T) {
	registerCommand := entity.SummonerRegisterCommand{
		PlayerID: "test_player_id",
		Monsters: []entity.Monster{
			{
				Name: "test_monster",
			},
		},
	}

	tests := []struct {
		name string

		command entity.SummonerRegisterCommand

		mockSummoner *entity.Summoner
		findError    error
		saveError    error

		expectedError bool
	}{
		{
			name: "ShouldRegisterSuccessfully",

			command: registerCommand,

			mockSummoner: nil,

			findError: nil,
			saveError: nil,

			expectedError: false,
		},
		{
			name: "ShouldReturnErrorWhenSummonerAlreadyExists",

			command: registerCommand,

			mockSummoner: &entity.Summoner{
				PlayerID: "test_player_id",
				Monsters: []entity.Monster{
					{
						Name: "test_monster",
					},
				},
			},
			findError: nil,
			saveError: nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenSummonerCreationFails",

			command: entity.SummonerRegisterCommand{
				PlayerID: "",
				Monsters: []entity.Monster{},
			},

			mockSummoner: nil,
			findError:    nil,
			saveError:    nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenFindSummonerFails",

			command: registerCommand,

			mockSummoner: nil,
			findError:    errors.New("find error"),
			saveError:    nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenSaveSummonerFails",

			command: registerCommand,

			mockSummoner: nil,
			findError:    nil,
			saveError:    errors.New("save error"),

			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGetter := &mockSummonerGetter{
				mockSummoner: tt.mockSummoner,
				findError:    tt.findError,
			}
			mockDomainService := NewSummonerService(mockGetter)

			mockSaver := &mockSummonerSaver{
				saveError: tt.saveError,
			}
			summonerRegisterService := NewSummonerRegisterService(mockSaver, mockDomainService)

			err := summonerRegisterService.Register(tt.command)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}

package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type mockSummonerGetter struct {
	mockSummoner *entity.Summoner
	findError    error
}

func (m *mockSummonerGetter) Find(playerID string) (*entity.Summoner, error) {
	return m.mockSummoner, m.findError
}

func TestSummonerServiceExists(t *testing.T) {
	summoner := entity.Summoner{
		PlayerID: "test_player_id",
		Monsters: []string{},
	}

	tests := []struct {
		name string

		mockSummoner *entity.Summoner
		findError    error

		expectedBool  bool
		expectedError bool
	}{
		{
			name: "TestSummonerServiceExistsShouldReturnFalseWithSummonerNotFound",

			mockSummoner: nil,
			findError:    nil,

			expectedBool:  false,
			expectedError: false,
		},
		{
			name: "SummonerServiceExistsShouldReturTrueWhenSummonerExists",

			mockSummoner: &summoner,
			findError:    nil,

			expectedBool:  true,
			expectedError: false,
		},
		{
			name: "SummonerServiceExistsShouldReturnFalseWhenRepositoryReturnsError",

			mockSummoner: nil,
			findError:    errors.New("repository error"),

			expectedBool:  false,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := &mockSummonerGetter{
				mockSummoner: tt.mockSummoner,
				findError:    tt.findError,
			}
			// Initialize the service
			service := NewSummonerService(mockRepository)

			exists, err := service.Exists("test_player_id")

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if exists != tt.expectedBool {
				t.Errorf("expected exists: %v, got: %v", tt.expectedBool, exists)
			}
		})
	}
}

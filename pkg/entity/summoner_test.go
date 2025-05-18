package entity

import "testing"

func TestNewSummoner(t *testing.T) {
	tests := []struct {
		name string

		playerID string

		expectedError bool
	}{
		{
			name: "NewSummonerShouldReturnSuccessfully",

			playerID: "test_player_id",

			expectedError: false,
		},
		{
			name: "NewSummonerShouldReturnErrorWhenPlayerIDIsEmpty",

			playerID: "",

			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			summoner, err := NewSummoner(tt.playerID)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if err == nil && summoner.PlayerID != tt.playerID {
				t.Errorf("expected player_id: %s, got: %s", tt.playerID, summoner.PlayerID)
			}
		})
	}
}

func TestAddMonster(t *testing.T) {
	tests := []struct {
		name string

		summoner Summoner
		monster  Monster
	}{
		{
			name: "AddMonsterShouldReturnSuccessfully",

			summoner: Summoner{
				PlayerID: "test_player_id",
				Monsters: []Monster{},
			},
			monster: Monster{
				Name: "test_monster",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.summoner.AddMonster(tt.monster)
			if len(tt.summoner.Monsters) != 1 {
				t.Errorf("expected 1 monster, got: %d", len(tt.summoner.Monsters))
			}
		})
	}
}

func TestSummonerEqual(t *testing.T) {
	tests := []struct {
		name string

		summoner1 *Summoner
		summoner2 *Summoner

		expectedEqual bool
	}{
		{
			name: "EqualSummonersShouldReturnTrue",

			summoner1: &Summoner{
				PlayerID: "test_player_id",
				Monsters: []Monster{},
			},
			summoner2: &Summoner{
				PlayerID: "test_player_id",
				Monsters: []Monster{},
			},

			expectedEqual: true,
		},
		{
			name: "DifferentPlayerIDsShouldReturnFalse",

			summoner1: &Summoner{
				PlayerID: "test_player_id",
				Monsters: []Monster{},
			},
			summoner2: &Summoner{
				PlayerID: "different_player_id",
				Monsters: []Monster{},
			},

			expectedEqual: false,
		},
		{
			name: "DifferentMonstersShouldReturnFalse",
			summoner1: &Summoner{
				PlayerID: "test_player_id",
				Monsters: []Monster{
					{
						Name: "monster1",
					},
				},
			},
			summoner2: &Summoner{
				PlayerID: "test_player_id",
				Monsters: []Monster{},
			},

			expectedEqual: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.summoner1.Equal(tt.summoner2)
			if result != tt.expectedEqual {
				t.Errorf("expected equal: %v, got: %v", tt.expectedEqual, result)
			}
		})
	}
}

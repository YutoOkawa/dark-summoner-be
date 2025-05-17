package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

func TestMonsterGetInfoServiceGetInfo(t *testing.T) {
	monster := entity.Monster{
		Name: "test_monster",
		ATK: entity.MonsterParameter{
			Parameter: 10,
		},
		DEF: entity.MonsterParameter{
			Parameter: 5,
		},
		HP: entity.MonsterParameter{
			Parameter: 20,
		},
		Rarity:    "test",
		RarityID:  0,
		Skill:     "test",
		Attribute: "test",
	}

	tests := []struct {
		name        string
		monsterName string

		mockMonster *entity.Monster
		findError   error

		expectedMonster *entity.Monster
		expectedError   bool
	}{
		{
			name:        "ShouldReturnGetInfoSuccessfully",
			monsterName: "test_monster",

			mockMonster: &monster,
			findError:   nil,

			expectedMonster: &monster,
			expectedError:   false,
		},
		{
			name:        "ShouldReturnErrorWhenMonsterDoesNotExist",
			monsterName: "non_existent_monster",

			mockMonster: nil,
			findError:   nil,

			expectedMonster: nil,
			expectedError:   true,
		},
		{
			name:        "ShouldReturnErrorWhenRepositoryReturnsError",
			monsterName: "test_monster",

			mockMonster: nil,
			findError:   errors.New("repository error"),

			expectedMonster: nil,
			expectedError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mockMonsterGetter{
				mockMonster: tt.mockMonster,
				findError:   tt.findError,
			}

			service := NewMonsterGetInfoService(&repo)
			monster, err := service.GetInfo(tt.monsterName)

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if monster != tt.expectedMonster {
				t.Errorf("expected monster: %v, got: %v", tt.expectedMonster, monster)
			}
		})
	}

}

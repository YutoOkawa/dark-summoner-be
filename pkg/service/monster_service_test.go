package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type mockMonsterGetter struct {
	mockMonster *entity.Monster
	findError   error
}

func (m *mockMonsterGetter) Find(name entity.MonsterName) (*entity.Monster, error) {
	return m.mockMonster, m.findError
}

func TestMonsterCheckServiceExists(t *testing.T) {
	monster := entity.Monster{
		Name:      entity.MonsterName("test_monster"),
		ATK:       entity.MonsterParameter(10),
		DEF:       entity.MonsterParameter(5),
		HP:        entity.MonsterParameter(100),
		Rarity:    "test",
		RarityID:  0,
		Skill:     "test",
		Attribute: "test",
	}

	tests := []struct {
		name string

		mockMonster *entity.Monster
		findError   error

		expectedBool  bool
		expectedError bool
	}{
		{
			name:          "ShouldReturnTrueWhenMonsterExists",
			mockMonster:   &monster,
			expectedBool:  true,
			expectedError: false,
		},
		{
			name:          "ShouldReturnFalseWhenMonsterDoesNotExist",
			mockMonster:   nil,
			expectedBool:  false,
			expectedError: false,
		},
		{
			name:        "ShouldReturnFalseWhenRepositoryReturnsError",
			mockMonster: &monster,
			findError:   errors.New("Getter error"),

			expectedBool:  false,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mockMonsterGetter{
				mockMonster: tt.mockMonster,
				findError:   tt.findError,
			}

			service := NewMonsterService(&repo)
			exists, err := service.Exists(monster)

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
			if exists != tt.expectedBool {
				t.Errorf("expected exists: %v, got: %v", tt.expectedBool, exists)
			}
		})
	}

}

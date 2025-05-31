package repository

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

func TestInMemoryMonsterRepositorySave(t *testing.T) {
	repo := NewInMemoryMonsterRepository()

	tests := []struct {
		name     string
		monster  entity.Monster
		expected error
	}{
		{
			name: "SaveMonsterSuccessfully",
			monster: entity.Monster{
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
			},
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Save(tt.monster)
			if err != nil && err.Error() != tt.expected.Error() {
				t.Errorf("expected error: %v, got: %v", tt.expected, err)
			}
		})
	}
}

func TestInMemoryMonsterRepositoryFind(t *testing.T) {
	repo := NewInMemoryMonsterRepository()
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
	repo.Save(monster)

	tests := []struct {
		name        string
		monsterName string
		expected    *entity.Monster
	}{
		{
			name: "FindExistingMonster",

			monsterName: "test_monster",

			expected: &monster,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := repo.Find(tt.monsterName)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result == nil || *result != *tt.expected {
				t.Errorf("expected: %v, got: %v", *tt.expected, result)
			}
		})
	}
}

func TestInMemoryMonsterRepositoryList(t *testing.T) {
	tests := []struct {
		name string

		monsters []entity.Monster

		expectedError error
	}{
		{
			name: "ShouldInMemoryMonsterRepositoryListSuccessfully",

			monsters: []entity.Monster{
				{
					Name: "monster1",
				},
				{
					Name: "monster2",
				},
			},

			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := InMemoryMonsterRepository{
				monsters: tt.monsters,
			}

			monsters, err := repo.List()
			if len(monsters) != len(tt.monsters) {
				t.Errorf("expected %d monsters, got %d", len(tt.monsters), len(monsters))
			}
			if !errors.Is(err, tt.expectedError) {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}

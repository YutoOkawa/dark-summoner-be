package repository

import (
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
				Name: entity.MonsterName{
					Name: "test_monster",
				},
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
		Name: entity.MonsterName{
			Name: "test_monster",
		},
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
		monsterName entity.MonsterName
		expected    *entity.Monster
	}{
		{
			name: "FindExistingMonster",
			monsterName: entity.MonsterName{
				Name: "test_monster",
			},
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

package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type mockMonsterGetterSaver struct {
	mockMonster *entity.Monster
	findError   error
	saveError   error
}

func (m *mockMonsterGetterSaver) Find(name entity.MonsterName) (*entity.Monster, error) {
	return m.mockMonster, m.findError
}

func (m *mockMonsterGetterSaver) Save(monster entity.Monster) error {
	return m.saveError
}

func TestMonsterRegisterService(t *testing.T) {
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
		saveError   error

		expectedError bool
	}{
		{
			name: "ShouldRegisterSuccessfully",

			mockMonster: nil,
			findError:   nil,
			saveError:   nil,

			expectedError: false,
		},
		{
			name: "ShouldReturnErrorWhenMonsterAlreadyExists",

			mockMonster: &monster,
			findError:   nil,
			saveError:   nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenFindFails",

			mockMonster: nil,
			findError:   errors.New("find error"),
			saveError:   nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenSaveFails",

			mockMonster: nil,
			findError:   nil,
			saveError:   errors.New("save error"),

			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mockMonsterGetterSaver{
				mockMonster: tt.mockMonster,
				findError:   tt.findError,
				saveError:   tt.saveError,
			}
			domainService := NewMonsterService(&repo)
			service := NewMonsterRegisterService(&repo, domainService)

			err := service.Register(monster)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}

}

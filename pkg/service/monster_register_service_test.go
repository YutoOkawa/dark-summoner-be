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
	monsterRegisterCommand := entity.MonsterRegisterCommand{
		Name:      "test_monster",
		ATK:       10,
		DEF:       5,
		HP:        100,
		Rarity:    "test",
		RarityID:  0,
		Skill:     "test",
		Attribute: "test",
	}

	tests := []struct {
		name    string
		command entity.MonsterRegisterCommand

		mockMonster *entity.Monster
		findError   error
		saveError   error

		expectedError bool
	}{
		{
			name:    "ShouldRegisterSuccessfully",
			command: monsterRegisterCommand,

			mockMonster: nil,
			findError:   nil,
			saveError:   nil,

			expectedError: false,
		},
		{
			name:    "ShouldReturnErrorWhenMonsterAlreadyExists",
			command: monsterRegisterCommand,

			mockMonster: &entity.Monster{
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
			findError: nil,
			saveError: nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenMonsterNameIsInvalid",
			command: entity.MonsterRegisterCommand{
				Name: "",
			},

			mockMonster: nil,
			findError:   nil,
			saveError:   nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenMonsterATKIsInvalid",
			command: entity.MonsterRegisterCommand{
				Name: "test_monster",
				ATK:  -1,
			},

			mockMonster: nil,
			findError:   nil,
			saveError:   nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenMonsterDEFIsInvalid",
			command: entity.MonsterRegisterCommand{
				Name: "test_monster",
				ATK:  10,
				DEF:  -1,
			},

			mockMonster: nil,
			findError:   nil,
			saveError:   nil,

			expectedError: true,
		},
		{
			name: "ShouldReturnErrorWhenMonsterHPIsInvalid",
			command: entity.MonsterRegisterCommand{
				Name: "test_monster",
				ATK:  10,
				DEF:  5,
				HP:   -1,
			},

			mockMonster: nil,
			findError:   nil,
			saveError:   nil,

			expectedError: true,
		},
		{
			name:    "ShouldReturnErrorWhenFindFails",
			command: monsterRegisterCommand,

			mockMonster: nil,
			findError:   errors.New("find error"),
			saveError:   nil,

			expectedError: true,
		},
		{
			name:    "ShouldReturnErrorWhenSaveFails",
			command: monsterRegisterCommand,

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

			err := service.Register(tt.command)
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}

}

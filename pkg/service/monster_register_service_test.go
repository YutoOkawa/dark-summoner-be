package service

import (
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
		name          string
		monster       *entity.Monster
		expectedError bool
	}{
		{
			name:          "ShouldRegisterSuccessfully",
			monster:       nil,
			expectedError: false,
		},
		{
			name:          "ShouldReturnErrorWhenMonsterAlreadyExists",
			monster:       &monster,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := mockMonsterGetterSaver{
				mockMonster: tt.monster,
				findError:   nil,
				saveError:   nil,
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

package service

import (
	"errors"
	"testing"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
)

type mockMonsterLister struct {
	monsters  []entity.Monster
	listError error
}

func (m *mockMonsterLister) List() ([]entity.Monster, error) {
	return m.monsters, m.listError
}

func TestMonsterListServiceList(t *testing.T) {
	tests := []struct {
		name          string
		monsters      []entity.Monster
		mockListError error

		expectedMonstersLength int
		expectedError          bool
	}{
		{
			name: "ShouldMonsterListServiceListSuccessfully",

			monsters: []entity.Monster{
				{
					Name: "monster1",
				},
				{
					Name: "monster2",
				},
			},
			mockListError: nil,

			expectedMonstersLength: 2,
			expectedError:          false,
		},
		{
			name:          "ShouldMonsterListServiceListReturnError",
			monsters:      []entity.Monster{},
			mockListError: errors.New("list error"),

			expectedMonstersLength: 0,
			expectedError:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mockMonsterLister{
				monsters:  tt.monsters,
				listError: tt.mockListError,
			}

			service := NewMonsterListService(repo)
			monsters, err := service.List()
			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err != nil)
			}
			if len(monsters) != tt.expectedMonstersLength {
				t.Errorf("expected %d monsters, got %d", tt.expectedMonstersLength, len(monsters))
			}
		})
	}
}

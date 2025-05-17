package entity

import "testing"

func TestNewMonsterName(t *testing.T) {
	expectedMonsterName := MonsterName{
		Name: "test_monster",
	}

	tests := []struct {
		name                string
		inputName           string
		expectedMonsterName *MonsterName
		err                 bool
	}{
		{
			name:                "NewMonsterNameShouldReturnSuccessfully",
			inputName:           "test_monster",
			expectedMonsterName: &expectedMonsterName,
			err:                 false,
		},
		{
			name:                "NewMonsterNameShouldReturnErrorWhenNameIsEmpty",
			inputName:           "",
			expectedMonsterName: nil,
			err:                 true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewMonsterName(tt.inputName)
			if (err != nil) != tt.err {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if result != nil && *result != *tt.expectedMonsterName {
				t.Errorf("expected: %v, got: %v", *tt.expectedMonsterName, *result)
			}
		})
	}
}

func TestNewMonsterParameter(t *testing.T) {
	expectedMonsterParameter := MonsterParameter{
		Parameter: 10,
	}

	tests := []struct {
		name                     string
		inputParameter           int
		expectedMonsterParameter *MonsterParameter
		err                      bool
	}{
		{
			name:                     "NewMonsterParameterShouldReturnSuccessfully",
			inputParameter:           10,
			expectedMonsterParameter: &expectedMonsterParameter,
			err:                      false,
		},
		{
			name:                     "NewMonsterParameterShouldReturnErrorWhenParameterIsNegative",
			inputParameter:           -1,
			expectedMonsterParameter: nil,
			err:                      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewMonsterParameter(tt.inputParameter)
			if (err != nil) != tt.err {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if result != nil && *result != *tt.expectedMonsterParameter {
				t.Errorf("expected: %v, got: %v", *tt.expectedMonsterParameter, *result)
			}
		})
	}
}

func TestNewMonster(t *testing.T) {
	expectedMonster := &Monster{
		Name: MonsterName{
			Name: "test_monster",
		},
		ATK: MonsterParameter{
			Parameter: 10,
		},
		DEF: MonsterParameter{
			Parameter: 5,
		},
		HP: MonsterParameter{
			Parameter: 20,
		},
		Rarity:    "test",
		RarityID:  0,
		Skill:     "test",
		Attribute: "test",
	}

	tests := []struct {
		name            string
		inputMonster    *Monster
		expectedMonster *Monster
		err             bool
	}{
		{
			name:            "NewMonsterShouldReturnSuccessfully",
			inputMonster:    expectedMonster,
			expectedMonster: expectedMonster,
			err:             false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewMonster(
				tt.inputMonster.Name,
				tt.inputMonster.ATK,
				tt.inputMonster.DEF,
				tt.inputMonster.HP,
				tt.inputMonster.Rarity,
				tt.inputMonster.RarityID,
				tt.inputMonster.Skill,
				tt.inputMonster.Attribute,
			)
			if (err != nil) != tt.err {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if result != nil && *result != *tt.expectedMonster {
				t.Errorf("expected: %v, got: %v", *tt.expectedMonster, *result)
			}
		})
	}
}

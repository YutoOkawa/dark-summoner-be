package entity

import "errors"

type MonsterName string

func NewMonsterName(name string) (*MonsterName, error) {
	if name == "" {
		return nil, errors.New("illegal argument: name cannot be empty")
	}
	monsterName := MonsterName(name)
	return &monsterName, nil
}

type MonsterParameter int

func NewMonsterParameter(param int) (*MonsterParameter, error) {
	if param < 0 {
		return nil, errors.New("illegal argument: parameter cannot be negative")
	}
	monsterParam := MonsterParameter(param)
	return &monsterParam, nil
}

type Monster struct {
	Name      MonsterName      `json:"name"`
	ATK       MonsterParameter `json:"atk"`
	DEF       MonsterParameter `json:"def"`
	HP        MonsterParameter `json:"hp"`
	Rarity    string           `json:"rarity"`
	RarityID  int              `json:"rarity_id"`
	Skill     string           `json:"skill"`
	Attribute string           `json:"attribute"`
}

func NewMonster(name MonsterName, atk, def, hp MonsterParameter, rarity string, rarityID int, skill, attribute string) (*Monster, error) {
	if name == "" {
		return nil, errors.New("illegal argument: name cannot be empty")
	}
	if atk < 0 {
		return nil, errors.New("illegal argument: atk cannot be negative")
	}
	if def < 0 {
		return nil, errors.New("illegal argument: def cannot be negative")
	}
	if hp < 0 {
		return nil, errors.New("illegal argument: hp cannot be negative")
	}
	if rarity == "" {
		return nil, errors.New("illegal argument: rarity cannot be empty")
	}
	if rarityID < 0 {
		return nil, errors.New("illegal argument: rarity_id cannot be negative")
	}
	if skill == "" {
		return nil, errors.New("illegal argument: skill cannot be empty")
	}
	if attribute == "" {
		return nil, errors.New("illegal argument: attribute cannot be empty")
	}

	return &Monster{
		Name:      name,
		ATK:       atk,
		DEF:       def,
		HP:        hp,
		Rarity:    rarity,
		RarityID:  rarityID,
		Skill:     skill,
		Attribute: attribute,
	}, nil
}

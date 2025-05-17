package entity

import "errors"

type MonsterParameter struct {
	Parameter int
}

func NewMonsterParameter(param int) (*MonsterParameter, error) {
	if param < 0 {
		return nil, errors.New("illegal argument: parameter cannot be negative")
	}
	return &MonsterParameter{
		Parameter: param,
	}, nil
}

type Monster struct {
	Name      string           `json:"name"`
	ATK       MonsterParameter `json:"atk"`
	DEF       MonsterParameter `json:"def"`
	HP        MonsterParameter `json:"hp"`
	Rarity    string           `json:"rarity"`
	RarityID  int              `json:"rarity_id"`
	Skill     string           `json:"skill"`
	Attribute string           `json:"attribute"`
}

func NewMonster(name string, atk, def, hp MonsterParameter, rarity string, rarityID int, skill, attribute string) (*Monster, error) {
	if name == "" {
		return nil, errors.New("illegal argument: name cannot be empty")
	}
	if atk == (MonsterParameter{}) {
		return nil, errors.New("illegal argument: atk cannot be empty")
	}
	if def == (MonsterParameter{}) {
		return nil, errors.New("illegal argument: def cannot be empty")
	}
	if hp == (MonsterParameter{}) {
		return nil, errors.New("illegal argument: hp cannot be empty")
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

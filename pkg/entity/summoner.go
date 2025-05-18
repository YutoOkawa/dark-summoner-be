package entity

import "errors"

type Summoner struct {
	PlayerID string    `json:"player_id"`
	Monsters []Monster `json:"monsters"`
}

func NewSummoner(playerID string) (*Summoner, error) {
	if playerID == "" {
		return nil, errors.New("illegal argument: player_id cannot be negative")
	}
	return &Summoner{
		PlayerID: playerID,
		Monsters: []Monster{},
	}, nil
}

func (s *Summoner) AddMonster(monster Monster) {
	s.Monsters = append(s.Monsters, monster)
}

func (s *Summoner) Equal(other *Summoner) bool {
	if s.PlayerID != other.PlayerID {
		return false
	}
	if len(s.Monsters) != len(other.Monsters) {
		return false
	}
	return true
}

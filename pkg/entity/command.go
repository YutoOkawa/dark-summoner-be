package entity

type MonsterRegisterCommand struct {
	Name      string
	ATK       int
	DEF       int
	HP        int
	Rarity    string
	RarityID  int
	Skill     string
	Attribute string
}

type SummonerRegisterCommand struct {
	PlayerID string
	Monsters []Monster
}

type SummonerAddMonsterCommand struct {
	PlayerID string
	Monster  Monster
}

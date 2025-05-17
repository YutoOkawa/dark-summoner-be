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

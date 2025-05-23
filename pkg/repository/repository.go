package repository

import "github.com/YutoOkawa/dark-summoner-be/pkg/entity"

type MonsterSaver interface {
	Save(monster entity.Monster) error
}

type MonsterGetter interface {
	Find(name string) (*entity.Monster, error)
}

type MonsterGetterSaver interface {
	MonsterSaver
	MonsterGetter
}

type SummonerSaver interface {
	Save(summoner entity.Summoner) error
}

type SummonerGetter interface {
	Find(playerID string) (*entity.Summoner, error)
}

type SummonerGetterSaver interface {
	SummonerSaver
	SummonerGetter
}

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

package handler

import (
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type MonsterListHandler struct {
	monsterListService service.MonsterListService
}

func NewMonsterListHandler(monsterListService service.MonsterListService) MonsterListHandler {
	return MonsterListHandler{
		monsterListService: monsterListService,
	}
}

func (l MonsterListHandler) ListHandlerFunc() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		monsters, err := l.monsterListService.List()
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(monsters)
	}
}

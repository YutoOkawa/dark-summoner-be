package handler

import (
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type GetInfoHandler struct {
	monsterGetInfoService service.MonsterGetInfoService
}

func NewGetInfoHandler(monsterGetInfoService service.MonsterGetInfoService) GetInfoHandler {
	return GetInfoHandler{
		monsterGetInfoService: monsterGetInfoService,
	}
}

func (g *GetInfoHandler) GetInfoHandlerFunc() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")

		monster, err := g.monsterGetInfoService.GetInfo(name)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if monster == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.JSON(monster)
	}
}

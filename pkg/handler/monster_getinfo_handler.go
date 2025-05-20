package handler

import (
	"net/url"

	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type MonsterGetInfoHandler struct {
	monsterGetInfoService service.MonsterGetInfoService
}

func NewGetInfoHandler(monsterGetInfoService service.MonsterGetInfoService) MonsterGetInfoHandler {
	return MonsterGetInfoHandler{
		monsterGetInfoService: monsterGetInfoService,
	}
}

func (g *MonsterGetInfoHandler) GetInfoHandlerFunc() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")

		unescapeName, err := url.PathUnescape(name)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		monster, err := g.monsterGetInfoService.GetInfo(unescapeName)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if monster == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.JSON(monster)
	}
}

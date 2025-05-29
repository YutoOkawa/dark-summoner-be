package handler

import (
	"net/url"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type SummonerAddMonsterHandler struct {
	summonerService           service.SummonerService
	summonerAddMonsterService service.SummonerAddMonsterService
	monsterGetInfoService     service.MonsterGetInfoService
}

func NewSummonerAddMonsterHandler(summonerService service.SummonerService, summonerAddMonsterService service.SummonerAddMonsterService, monsterGetInfoService service.MonsterGetInfoService) SummonerAddMonsterHandler {
	return SummonerAddMonsterHandler{
		summonerService:           summonerService,
		summonerAddMonsterService: summonerAddMonsterService,
		monsterGetInfoService:     monsterGetInfoService,
	}
}

func (h *SummonerAddMonsterHandler) AddMonsterHandlerFunc() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		playerID := c.Params("player_id")
		if playerID == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		monsterName := c.Params("monster_name")
		if monsterName == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		unescapedMonsterName, err := url.PathUnescape(monsterName)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		exists, err := h.summonerService.Exists(playerID)
		if err != nil {
			return c.SendString(err.Error())
		}

		if !exists {
			return c.SendStatus(fiber.StatusNotFound)
		}

		monster, err := h.monsterGetInfoService.GetInfo(unescapedMonsterName)
		if err != nil {
			return c.SendString(err.Error())
		}

		if monster == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		command := entity.SummonerAddMonsterCommand{
			PlayerID: playerID,
			Monster:  *monster,
		}

		err = h.summonerAddMonsterService.AddMonster(command)
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

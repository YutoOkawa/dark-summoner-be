package handler

import (
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type SummonerGetInfoHandler struct {
	summonerGetinfoService service.SummonerGetInfoService
}

func NewSummonerGetInfoHandler(summonerGetinfoService service.SummonerGetInfoService) SummonerGetInfoHandler {
	return SummonerGetInfoHandler{
		summonerGetinfoService: summonerGetinfoService,
	}
}

func (r *SummonerGetInfoHandler) GetInfoHandlerFunc() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		playerID := c.Params("player_id")
		if playerID == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		summoner, err := r.summonerGetinfoService.GetInfo(playerID)
		if err != nil {
			return c.SendString(err.Error())
		}
		if summoner == nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.JSON(summoner)
	}
}

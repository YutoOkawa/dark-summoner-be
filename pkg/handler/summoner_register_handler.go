package handler

import (
	"encoding/json"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type SummonerRegisterRequestParam struct {
	PlayerID string `json:"player_id"`
}

type SummonerRegisterHandler struct {
	summonerRegisterService service.SummonerRegisterService
}

func NewSummonerRegisterHandler(summonerRegisterService service.SummonerRegisterService) SummonerRegisterHandler {
	return SummonerRegisterHandler{
		summonerRegisterService: summonerRegisterService,
	}
}

func (r *SummonerRegisterHandler) RegisterHandlerFunc() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := c.Body()

		var requestParam SummonerRegisterRequestParam
		if err := json.Unmarshal(body, &requestParam); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		command := entity.SummonerRegisterCommand{
			PlayerID: requestParam.PlayerID,
			Monsters: []entity.Monster{},
		}

		err := r.summonerRegisterService.Register(command)
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

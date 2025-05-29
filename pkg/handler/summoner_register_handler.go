package handler

import (
	"encoding/json"
	"net/url"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type SummonerRegisterRequestParam struct {
	PlayerID string   `json:"player_id"`
	Monsters []string `json:"monsters"`
}

type SummonerRegisterHandler struct {
	summonerService         *service.SummonerService
	summonerRegisterService service.SummonerRegisterService
	monsterGetInfoService   service.MonsterGetInfoService
}

func NewSummonerRegisterHandler(summonerService *service.SummonerService, summonerRegisterService service.SummonerRegisterService, monsterGetInfoService service.MonsterGetInfoService) SummonerRegisterHandler {
	return SummonerRegisterHandler{
		summonerService:         summonerService,
		summonerRegisterService: summonerRegisterService,
		monsterGetInfoService:   monsterGetInfoService,
	}
}

func (r *SummonerRegisterHandler) RegisterHandlerFunc() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := c.Body()

		var requestParam SummonerRegisterRequestParam
		if err := json.Unmarshal(body, &requestParam); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		exists, err := r.summonerService.Exists(requestParam.PlayerID)
		if err != nil {
			return c.SendString(err.Error())
		}
		if exists {
			return c.SendStatus(fiber.StatusConflict)
		}

		monsters := make([]entity.Monster, 0, len(requestParam.Monsters))
		for _, monsterName := range requestParam.Monsters {
			unescapedMonsterName, err := url.PathUnescape(monsterName)
			if err != nil {
				return c.SendStatus(fiber.StatusBadRequest)
			}

			monster, err := r.monsterGetInfoService.GetInfo(unescapedMonsterName)
			if err != nil {
				return c.SendString(err.Error())
			}
			if monster == nil {
				return c.SendStatus(fiber.StatusNotFound)
			}
			monsters = append(monsters, *monster)
		}

		command := entity.SummonerRegisterCommand{
			PlayerID: requestParam.PlayerID,
			Monsters: monsters,
		}

		err = r.summonerRegisterService.Register(command)
		if err != nil {
			return c.SendString(err.Error())
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

package handler

import (
	"encoding/json"

	"github.com/YutoOkawa/dark-summoner-be/pkg/entity"
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type MonsterRegisterRequestParam struct {
	Name      string `json:"name"`
	ATK       int    `json:"atk"`
	DEF       int    `json:"def"`
	HP        int    `json:"hp"`
	Rarity    string `json:"rarity"`
	RarityID  int    `json:"rarity_id"`
	Skill     string `json:"skill"`
	Attribute string `json:"attribute"`
}

type MonsterRegisterHandler struct {
	monsterRegisterService service.MonsterRegisterService
}

func NewRegisterHandler(monsterRegisterService service.MonsterRegisterService) MonsterRegisterHandler {
	return MonsterRegisterHandler{
		monsterRegisterService: monsterRegisterService,
	}
}

func (r *MonsterRegisterHandler) RegisterHandlerFunc() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		body := c.Body()

		var requestParam MonsterRegisterRequestParam
		if err := json.Unmarshal(body, &requestParam); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		command := entity.MonsterRegisterCommand{
			Name:      requestParam.Name,
			ATK:       requestParam.ATK,
			DEF:       requestParam.DEF,
			HP:        requestParam.HP,
			Rarity:    requestParam.Rarity,
			RarityID:  requestParam.RarityID,
			Skill:     requestParam.Skill,
			Attribute: requestParam.Attribute,
		}

		err := r.monsterRegisterService.Register(command)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/YutoOkawa/dark-summoner-be/pkg/handler"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
	"github.com/YutoOkawa/dark-summoner-be/pkg/server"
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
)

func main() {
	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer ctxCancel()

	server := server.NewServer(":8080", 0)

	monsterRepository := repository.NewInMemoryMonsterRepository()
	summonerRepository := repository.NewInMemorySummonerRepository()

	monsterService := service.NewMonsterService(monsterRepository)
	monsterGetInfoService := service.NewMonsterGetInfoService(monsterRepository)
	monsterRegisterService := service.NewMonsterRegisterService(monsterRepository, monsterService)

	summonerService := service.NewSummonerService(summonerRepository)
	summonerRegisterService := service.NewSummonerRegisterService(summonerRepository, summonerService)

	monsterRegisterHandler := handler.NewMonsterRegisterHandler(monsterRegisterService)
	monsterGetInfoHandler := handler.NewMonsterGetInfoHandler(monsterGetInfoService)

	summonerRegisterHandler := handler.NewSummonerRegisterHandler(summonerService, summonerRegisterService, monsterGetInfoService)
	summonerGetInfoHandler := handler.NewSummonerGetInfoHandler(service.NewSummonerGetInfoService(summonerRepository))

	server.App.Get("/healthz", handler.HealthZHandler)
	v1 := server.App.Group("/v1")
	{
		monster := v1.Group("/monster")
		{
			monster.Post("/", monsterRegisterHandler.RegisterHandlerFunc())
			monster.Get("/:name", monsterGetInfoHandler.GetInfoHandlerFunc())
		}

		summoner := v1.Group("/summoner")
		{
			summoner.Post("/", summonerRegisterHandler.RegisterHandlerFunc())
			summoner.Get("/:player_id", summonerGetInfoHandler.GetInfoHandlerFunc())
		}
	}

	srvErrCh := make(chan error, 1)
	go server.Start(srvErrCh)

	for {
		select {
		case err := <-srvErrCh:
			panic(err)
		case <-ctx.Done():
			if err := server.Shutdown(); err != nil {
				panic(err)
			}
			return
		}
	}
}

package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/YutoOkawa/dark-summoner-be/pkg/config"
	"github.com/YutoOkawa/dark-summoner-be/pkg/handler"
	"github.com/YutoOkawa/dark-summoner-be/pkg/repository"
	"github.com/YutoOkawa/dark-summoner-be/pkg/server"
	"github.com/YutoOkawa/dark-summoner-be/pkg/service"
)

var defaultConfigFilePath = "/etc/config/dark-summoner-be/config.yaml"

func main() {
	configFilePath := flag.String("config", defaultConfigFilePath, "Path to the configuration file")
	flag.Parse()

	config, err := config.LoadConfigFile(*configFilePath)
	if err != nil {
		panic(err)
	}

	ctx, ctxCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer ctxCancel()

	server := server.NewServer(config.Port, 0)

	monsterRepository := repository.NewInMemoryMonsterRepository()
	monsterRepository.LoadJSONFile(config.MonsterFilePath)
	summonerRepository := repository.NewInMemorySummonerRepository()
	summonerRepository.LoadJSONFile(config.SummonerFilePath)

	monsterService := service.NewMonsterService(monsterRepository)
	monsterGetInfoService := service.NewMonsterGetInfoService(monsterRepository)
	monsterListService := service.NewMonsterListService(monsterRepository)
	monsterRegisterService := service.NewMonsterRegisterService(monsterRepository, monsterService)

	summonerService := service.NewSummonerService(summonerRepository)
	summonerRegisterService := service.NewSummonerRegisterService(summonerRepository, summonerService)
	summonerAddMonsterService := service.NewSummonerAddMonsterService(summonerRepository)

	monsterRegisterHandler := handler.NewMonsterRegisterHandler(monsterRegisterService)
	monsterGetInfoHandler := handler.NewMonsterGetInfoHandler(monsterGetInfoService)
	monsterListHandler := handler.NewMonsterListHandler(monsterListService)

	summonerRegisterHandler := handler.NewSummonerRegisterHandler(summonerService, summonerRegisterService, monsterGetInfoService)
	summonerGetInfoHandler := handler.NewSummonerGetInfoHandler(service.NewSummonerGetInfoService(summonerRepository))
	summonerAddMonsterHandler := handler.NewSummonerAddMonsterHandler(*summonerService, summonerAddMonsterService, monsterGetInfoService)

	server.App.Get("/healthz", handler.HealthZHandler)
	v1 := server.App.Group("/v1")
	{
		monster := v1.Group("/monster")
		{
			monster.Post("/", monsterRegisterHandler.RegisterHandlerFunc())
			monster.Get("/:name", monsterGetInfoHandler.GetInfoHandlerFunc())
		}

		monsters := v1.Group("/monsters")
		{
			monsters.Get("/", monsterListHandler.ListHandlerFunc())
		}

		summoner := v1.Group("/summoner")
		{
			summoner.Post("/", summonerRegisterHandler.RegisterHandlerFunc())
			summoner.Get("/:player_id", summonerGetInfoHandler.GetInfoHandlerFunc())
			summoner.Post("/:player_id/monster/:monster_name", summonerAddMonsterHandler.AddMonsterHandlerFunc())
		}
	}

	srvErrCh := make(chan error, 1)
	go server.Start(srvErrCh)

	for {
		select {
		case err := <-srvErrCh:
			panic(err)
		case <-ctx.Done():
			if err := summonerRepository.SaveJSONFile(config.SummonerFilePath); err != nil {
				panic(err)
			}
			if err := monsterRepository.SaveJSONFile(config.MonsterFilePath); err != nil {
				panic(err)
			}
			if err := server.Shutdown(); err != nil {
				panic(err)
			}
			return
		}
	}
}

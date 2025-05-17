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

	repository := repository.NewInMemoryMonsterRepository()

	monsterService := service.NewMonsterService(repository)
	monsterGetInfoService := service.NewMonsterGetInfoService(repository)
	monsterRegisterService := service.NewMonsterRegisterService(repository, monsterService)

	registerHandler := handler.NewRegisterHandler(monsterRegisterService)
	getInfoHandler := handler.NewGetInfoHandler(monsterGetInfoService)

	server.App.Get("/healthz", handler.HealthZHandler)
	v1 := server.App.Group("/v1")
	{
		register := v1.Group("/register")
		{
			register.Post("/", registerHandler.RegisterHandlerFunc())
		}

		getInfo := v1.Group("/getInfo")
		{
			getInfo.Get("/:name", getInfoHandler.GetInfoHandlerFunc())
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

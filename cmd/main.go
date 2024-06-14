package main

import (
	_ "github.com/lib/pq"
	"github.com/sergkim13/short-link-backend-go/configs"
	"github.com/sergkim13/short-link-backend-go/pkg/handler"
	"github.com/sergkim13/short-link-backend-go/pkg/repository"
	"github.com/sergkim13/short-link-backend-go/pkg/server"
	"github.com/sergkim13/short-link-backend-go/pkg/service"
	"github.com/sirupsen/logrus"
)


func main() {
    configs.InitConfig()

    db, err := repository.NewPostgresDB(configs.EnvConfig)
    if err != nil {
        logrus.Fatalf("failed to initialize db: %s", err.Error())
    }

    repos := repository.NewRepository(db)
    services := service.NewService(repos)
    handlers := handler.NewHandler(services)

    srv := new(server.Server)
    if err := srv.Run(configs.EnvConfig.Port, handlers.InitRoutes()); err != nil {
        logrus.Fatalf("error occurred while running http server: %s", err.Error())
    }
}

package main

import (
	"github.com/sergkim13/short-link-backend-go/pkg/handler"
	"github.com/sergkim13/short-link-backend-go/pkg/server"
	"github.com/sirupsen/logrus"
)


func main() {
    handlers := new(handler.Handler)
    srv := new(server.Server)
    if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
        logrus.Fatalf("error occured while running http server: %s", err.Error())
    }
}

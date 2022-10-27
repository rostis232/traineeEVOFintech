package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rostis232/traineeEVOFintech"
	"github.com/rostis232/traineeEVOFintech/internal/config"
	"github.com/rostis232/traineeEVOFintech/pkg/handler"
	"github.com/rostis232/traineeEVOFintech/pkg/repository"
	"github.com/rostis232/traineeEVOFintech/pkg/service"
	"log"
)

func main() {
	fmt.Printf("Starting application on port %s\n", config.PortNumber)

	db, err := repository.NewPostgresDB(config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(traineeEVOFintech.Server)
	if err := srv.Run(config.PortNumber, handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

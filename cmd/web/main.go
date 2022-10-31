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

// @title Trainee EVO Fintech project (Transactions App API)
// @version 1.0
// @description API Server for Transactions Application

// @contact.name Rostyslav Pylypiv
// @contact.email rostislav.pylypiv@gmail.com

// @host localhost:8000
// @BasePath /

func main() {
	run()
}

func run() error {
	fmt.Printf("Starting application on port %s\n", config.PortNumber)

	db, err := repository.NewPostgresDB(config.DBConfig)
	if err != nil {
		log.Fatal(err)
		return err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(traineeEVOFintech.Server)
	if err := srv.Run(config.PortNumber, handlers.InitRoutes()); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

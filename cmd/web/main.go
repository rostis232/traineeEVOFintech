package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/rostis232/traineeEVOFintech"
	"github.com/rostis232/traineeEVOFintech/internal/config"
	"github.com/rostis232/traineeEVOFintech/pkg/handler"
	"github.com/rostis232/traineeEVOFintech/pkg/repository"
	"github.com/rostis232/traineeEVOFintech/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Trainee EVO Fintech project (Transactions App API)
// @version 1.0
// @description API Server for Transactions Application

// @contact.name Rostyslav Pylypiv
// @contact.email rostislav.pylypiv@gmail.com

// @host localhost:8000
// @BasePath /

func main() {
	db, err := repository.NewPostgresDB(config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(traineeEVOFintech.Server)

	go func() {
		if err := srv.Run(config.PortNumber, handlers.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("Starting application on port %s\n", config.PortNumber)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("Application shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Error occured on server shuttimg down: %v", err)
	}
	if err := db.Close(); err != nil {
		log.Printf("Error occured on DB connection close: %v", err)
	}
}

package main

import (
	"fmt"
	"github.com/rostis232/traineeEVOFintech"
	"github.com/rostis232/traineeEVOFintech/internal/config"
	"github.com/rostis232/traineeEVOFintech/pkg/handler"
	"log"
)

func main() {
	fmt.Printf("Starting application on port %s\n", config.PortNumber)

	handlers := new(handler.Handler)

	srv := new(traineeEVOFintech.Server)
	if err := srv.Run(config.PortNumber, handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

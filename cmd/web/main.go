package main

import (
	"fmt"
	"log"
	"net/http"
	"traineeEVOFintech/internal/config"
)

func main() {
	fmt.Printf("Starting application on port %s\n", config.PortNumber)

	srv := &http.Server{
		Addr:    config.PortNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
	"traineeEVOFintech/internal/config"
)

func main() {
	srv := &http.Server{
		Addr: config.PortNumber,
	}
	
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

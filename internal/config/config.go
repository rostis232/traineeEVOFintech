package config

import "github.com/rostis232/traineeEVOFintech/pkg/repository"

const PortNumber = "8000"

var DBConfig = repository.DBConfig{
	Host:     "db", //use 'db' if you use docker, and 'localhost' without docker
	Port:     "5432",
	Username: "postgres",
	Password: "qwerty",
	DBName:   "postgres",
	SSLMode:  "disable",
}

package config

import "github.com/rostis232/traineeEVOFintech/pkg/repository"

const PortNumber = "8000"
const OnProduction = false

var DBConfig = repository.DBConfig{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "qwerty",
	DBName:   "postgres",
	SSLMode:  "disable",
}

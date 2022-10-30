package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	for _, arg := range os.Args[1:] {
		if arg == "-m" || arg == "--migrate" {
			log.Println("Make Migrations...")
		}
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
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

	if err := doMigrate(); err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func doMigrate() error {
	log.Println("Make Migrations...")
	m, err := migrate.New("file://./schema", "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	} else {
		if err := m.Up(); err != nil {
			return err
		} else {
			log.Println("Migrations are successful")
			return nil
		}
	}
}

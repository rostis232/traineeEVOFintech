package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	if len(os.Args) > 1 {
		if os.Args[1] == "-m" || os.Args[1] == "--migrate" {
			doMigrate()
		}
	} else {
		mgr := ""
		log.Println("You don`t make migration. Want to do it? [y - yes/n - no]")
		for {
			fmt.Scan(&mgr)
			if mgr == "y" || mgr == "n" {
				break
			} else {
				log.Println("Wrong answer. Try again.")
				log.Println("You don`t make migration. Want to do it? [y - yes/n - no]")
			}
		}
		switch mgr {
		case "y":
			doMigrate()
		case "n":
			log.Println("Skip migration")
		default:

		}
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func doMigrate() {
	log.Println("Make Migrations...")
	m, err := migrate.New("file://./schema", "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Println(err)
	} else {
		if err := m.Up(); err != nil {
			log.Printf("Migrations error: '%s'\n", err)
		} else {
			log.Println("Migrations are successful")
		}
	}
}

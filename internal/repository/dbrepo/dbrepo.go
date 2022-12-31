package dbrepo

import (
	"database/sql"

	"github.com/Charik-Goyal/bookings/internal/config"
	"github.com/Charik-Goyal/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	// Because of this NewPostgresRepo we are getting postgresDBRepo as reciever in interface DatabaseRepo
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

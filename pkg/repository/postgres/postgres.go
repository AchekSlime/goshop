package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"goshop/pkg/repository"
)

const (
	UsersTable    = "users"
	ProductsTable = "products"
)

func NewPostgresDb(cfg repository.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DbName,
	))
	if err != nil {
		return nil, fmt.Errorf("can't connect to postgres db")
	}
	return db, nil
}

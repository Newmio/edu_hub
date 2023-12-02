package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	SslMode  string
}

func Init(c Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.User, c.DbName, c.Password, c.SslMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

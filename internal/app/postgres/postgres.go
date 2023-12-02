package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func OpenDb()(*sqlx.DB, error) {
	viper.AddConfigPath("internal/app/configs")
	viper.SetConfigName("postgres")
	err := viper.ReadInConfig()
	if err != nil{
		return nil, err
	}

	db, err := initDb(Config{
		Host: viper.GetString("host"),
		Port: viper.GetString("port"),
		User: viper.GetString("user"),
		DbName: viper.GetString("dbName"),
		Password: viper.GetString("password"),
		SslMode: viper.GetString("sslMode"),
	})
	if err != nil{
		return nil, err
	}

	return db, nil
}

type Config struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
	SslMode  string
}

func initDb(c Config) (*sqlx.DB, error) {
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

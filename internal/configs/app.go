package configs

import (
	"ed/internal/domain/adapter/http"
	"ed/internal/domain/repository"
	"ed/internal/domain/service"
	"ed/pkg/db/postgres"
	"ed/pkg/util"

	"github.com/gin-gonic/gin"
)

func InitEngine() (*gin.Engine, error) {
	database, err := postgres.Init(postgres.Config{
		Host:     "localhost",
		Port:     "5436",
		User:     "postgres",
		DbName:   "postgres",
		Password: "qwerty",
		SslMode:  "disable",
	})
	if err != nil {
		return nil, util.ErrTrace(err, util.Trace())
	}

	repo := repository.NewRepo(database)
	serv := service.NewService(repo)
	httpHand := http.NewHandler(serv)

	r := gin.Default()
	r = httpHand.InitHttpRoutes(r)

	return r, nil
}

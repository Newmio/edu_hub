package app

import (
	"ed"
	"ed/internal/app/postgres"
	"ed/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func InitEngine() (*gin.Engine, error) {
	database, err := postgres.OpenDb()
	if err != nil {
		return nil, ed.ErrTrace(err, ed.Trace())
	}

	userRepo := user.NewUserRepo(database)
	userService := user.NewUserService(userRepo)
	userHand := user.NewHandler(userService)

	r := gin.Default()

	r = userHand.InitUserRoutes(r)

	return r, nil
}

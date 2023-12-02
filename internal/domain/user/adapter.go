package user

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

type handler struct {
	s IUserService
}

func NewHandler(s IUserService) *handler {
	return &handler{s: s}
}

func (h *handler) InitUserRoutes(r *gin.Engine)*gin.Engine{
	r.POST("/createAccount", h.CreateAccountRout)
	return r
}

func (h *handler) CreateAccountRout(c *gin.Context) {
	var acc Account

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = json.Unmarshal(body, &acc)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	err = h.s.CreateAccount(&acc)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}

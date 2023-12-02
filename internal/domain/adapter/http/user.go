package http

import (
	"ed/internal/domain/model"
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
)

//func (h *handler)

func (h *handler) CreateAccountRout(c *gin.Context){
	var acc model.Account

	body, err := io.ReadAll(c.Request.Body)
	if err != nil{
		c.JSON(400, err.Error())
		return
	}

	err = json.Unmarshal(body, &acc)
	if err != nil{
		c.JSON(500, err.Error())
		return
	}

	err = h.s.CreateAccount(&acc)
	if err != nil{
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}
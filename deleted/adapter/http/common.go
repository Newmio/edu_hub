package http

import "github.com/gin-gonic/gin"

func (h *handler) InitRout(c *gin.Context){
	err := h.s.InitTables()
	if err != nil{
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, "ok")
}
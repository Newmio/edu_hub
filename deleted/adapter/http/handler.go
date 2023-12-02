package http

import (
	"ed/internal/domain/service"

	"github.com/gin-gonic/gin"
)

type handler struct{
	s *service.Service
}

func NewHandler(s *service.Service)*handler{
	return &handler{s: s}
}

func (h *handler) InitHttpRoutes(r *gin.Engine)*gin.Engine{
	r.GET("/init", h.InitRout)
	r.POST("/acc", h.CreateAccountRout)

	return r
}
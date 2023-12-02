package lesson

import "github.com/gin-gonic/gin"

type handler struct{
	s ILessonService
}

func NewHandler(s ILessonService)*handler{
	return &handler{s: s}
}

func (h *handler) InitLessonRoutes(r *gin.Engine)*gin.Engine{
	return r
}
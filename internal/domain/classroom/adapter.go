package classroom

import "github.com/gin-gonic/gin"

type handler struct{
	s IClassroomService
}

func NewHandler(s IClassroomService)*handler{
	return &handler{s: s}
}

func (h *handler) InitClassroomRoutes(r *gin.Engine)*gin.Engine{
	return r
}
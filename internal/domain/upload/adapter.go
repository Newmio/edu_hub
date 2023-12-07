package upload

import (
	"ed/internal/domain/request"

	"github.com/gin-gonic/gin"
)

type handler struct{
	s IUploadService
	request request.IRequestService
}

func NewHandler(s IUploadService, request request.IRequestService)*handler{
	return &handler{s: s, request: request}
}

func (h *handler) InitUploadRoutes(r *gin.Engine)*gin.Engine{

	r.GET("ws/upload", h.uploadRout)

	return r
}

func (h *handler) uploadRout(c *gin.Context){
	conn, err := h.request.Connect(c.Writer, c.Request, c.Request.Header)
	if err != nil{

	}
	defer conn.Close()
}
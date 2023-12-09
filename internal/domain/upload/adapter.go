package upload

import (
	"ed/internal/domain/logger"
	"ed/internal/domain/request"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type handler struct {
	s       IUploadService
	request request.IRequestService
	logger  logger.ILoggerService
}

func NewHandler(s IUploadService, request request.IRequestService, logger logger.ILoggerService) *handler {
	return &handler{s: s, request: request, logger: logger}
}

func (h *handler) InitUploadRoutes(r *gin.Engine) *gin.Engine {

	r.GET("ws/upload", h.uploadRout)

	return r
}

func (h *handler) uploadRout(c *gin.Context) {
	var data FileData
	log := h.logger.InitLog(c)

	conn, err := h.request.WsConnect(c.Writer, c.Request, c.Request.Header)
	if err != nil {
		h.logger.WsErrorResponse(log, err)
		return
	}
	defer conn.Close()

	file, err := h.s.CreateFile()
	if err != nil {
		h.logger.WsErrorResponse(log, err)
	}

	for {
		msg, err := h.request.WsReadText(conn)
		if err != nil {
			h.logger.WsErrorResponse(log, err)
			return
		}
		log.Body_req += "\n" + string(msg)

		if msg == nil {
			continue
		}

		err = json.Unmarshal(msg, &data)
		if err != nil {
			h.logger.WsErrorResponse(log, err)
			return
		}

		err = h.s.AppendChank(file, &data)
		if err != nil {
			h.logger.WsErrorResponse(log, err)
		}

		err = conn.WriteMessage(websocket.TextMessage, h.logger.WsDefaultResponse(log))
		if err != nil {
			h.logger.WsErrorResponse(log, err)
			return
		}

		if data.Last {
			return
		}
	}
}

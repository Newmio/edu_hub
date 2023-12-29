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

	r.GET("/ws/upload", h.uploadRout)
	r.GET("/ws/unload", h.unloadRout)

	return r
}

// TODO доделать загрузку файла
func (h *handler) unloadRout(c *gin.Context) {
	//var chanks [][]byte
	var data FileSettings
	log := h.logger.InitLog(c)

	conn, err := h.request.WsConnect(c.Writer, c.Request, c.Request.Header)
	if err != nil {
		h.logger.WsErrorResponse(log, err)
		return
	}

	for {

		msg, err := h.request.WsReadText(conn)
		if err != nil {
			h.logger.WsErrorResponse(log, err)
		}
		log.Body_req += "\n" + string(msg)

		err = json.Unmarshal(msg, &data)
		if err != nil {
			h.logger.WsErrorResponse(log, err)
			return
		}

		// chanks, err = h.s.ReadLastChanks(&data)
		// if err != nil {
		// 	h.logger.WsErrorResponse(log, err)
		// 	return
		// }
	}
}

func (h *handler) uploadRout(c *gin.Context) {
	var data FileSettings
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
	defer file.Close()

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			h.logger.WsErrorResponse(log, err)
			return
		}

		switch t {

		case websocket.BinaryMessage:
			err = h.s.AppendChank(file, msg)
			if err != nil {
				h.logger.WsErrorResponse(log, err)
				return
			}

		case websocket.TextMessage:
			log.Body_req += "\n" + string(msg)
			err := json.Unmarshal(msg, &data)
			if err != nil {
				h.logger.WsErrorResponse(log, err)
				return
			}

			err = h.s.UpdateFile(file, &data)
			if err != nil {
				h.logger.WsErrorResponse(log, err)
				return
			}

			h.logger.WsDefaultResponse(log)
			return
		}

		err = conn.WriteMessage(websocket.TextMessage, h.logger.WsDefaultResponse(log))
		if err != nil {
			h.logger.WsErrorResponse(log, err)
			return
		}
	}
}

package request

import (
	"bytes"
	"ed"
	"ed/internal/domain/logger"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type IRequestService interface {
	WsConnect(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error)
	WsReadText(conn *websocket.Conn) ([]byte, error)
	WsReadBinary(conn *websocket.Conn)([]byte, error)
}

type requestService struct {
	r logger.ILoggerService
}

func NewRequestService(r logger.ILoggerService) *requestService {
	return &requestService{r: r}
}

func (s *requestService) WsConnect(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, nil)
}

func (s *requestService) WsReadText(conn *websocket.Conn) ([]byte, error) {
	t, msg, err := conn.ReadMessage()
	if err != nil {
		return nil, ed.ErrTrace(err, ed.Trace())
	}

	switch t {
	case websocket.TextMessage:
		return msg, nil
	}

	return nil, nil
}

func (s *requestService) WsReadBinary(conn *websocket.Conn)([]byte, error){
	t, msg, err := conn.ReadMessage()
	if err != nil {
		return nil, ed.ErrTrace(err, ed.Trace())
	}

	switch t {
	case websocket.BinaryMessage:
		return msg, nil
	}

	return nil, nil
}

// ==================== HTTP ====================
func (s *requestService) HttpRequest(param Param) (*http.Response, error) {
	var body []byte

	log := logger.Log{
		Url:    param.Url,
		Method: param.Method,
	}

	client := &http.Client{}

	if param.BodyType == "JSON" {

		b, err := json.Marshal(param.Body)
		if err != nil {
			return nil, ed.ErrTrace(err, ed.Trace())
		}
		body = b

	} else if param.BodyType == "XML" {

		b, err := xml.Marshal(param.Body)
		if err != nil {
			return nil, ed.ErrTrace(err, ed.Trace())
		}
		body = b

	} else {
		body = nil
	}

	log.Body_req = string(body)
	log.Date_start = time.Now()

	req, err := http.NewRequest(param.Method, param.Url, bytes.NewBuffer(body))
	if err != nil {
		return nil, ed.ErrTrace(err, ed.Trace())
	}

	for key, value := range param.Headers {
		log.Headers_req += key + ": " + value.(string)
		req.Header.Set(key, value.(string))
	}

	resp, err := client.Do(req)
	if err != nil {
		if resp != nil {
			return nil, ed.ErrTrace(err, ed.Trace())
		}

		log.Status = 404
		log.Date_stop = time.Now()
		return &http.Response{StatusCode: 404}, ed.ErrTrace(err, ed.Trace())
	}

	log.Date_stop = time.Now()
	log.Status = resp.StatusCode

	bodyResp, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, ed.ErrTrace(err, ed.Trace())
	}
	log.Body_resp = string(bodyResp)

	for key, values := range resp.Header {
		log.Headers_resp = fmt.Sprintf("%s: %v", key, values)
	}

	log.Success = true

	if param.CreateLog {
		s.r.HttpDefaultResponse(&log)
	}

	return resp, nil
}

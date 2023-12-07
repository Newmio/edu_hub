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
	Connect(w http.ResponseWriter, r *http.Request, h http.Header)(*websocket.Conn, error)
}

type requestService struct {
	r logger.ILoggerService
}

func NewLoggerService(r logger.ILoggerService) *requestService {
	return &requestService{r: r}
}

func (s *requestService) Connect(w http.ResponseWriter, r *http.Request, h http.Header) (*websocket.Conn, error) {
	return upgrader.Upgrade(w, r, h)
}

func (s *requestService) HttpRequest(param Param) (*http.Response, error) {
	var body []byte

	log := logger.Log{
		Url:    param.Url,
		Method: param.Method,
		Type:   "http",
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
		s.r.LoggerRun(&log, "")
	}

	return resp, nil
}

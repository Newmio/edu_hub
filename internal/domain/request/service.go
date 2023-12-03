package request

import (
	"bytes"
	"ed"
	"ed/internal/domain/logger"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type IRequestService interface {
}

type requestService struct {
	r logger.ILoggerService
}

func NewLoggerService(r logger.ILoggerService) *requestService {
	return &requestService{r: r}
}

func (s *requestService) RequestServer(param Param) (*http.Response, error) {
	var body []byte
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

	req, err := http.NewRequest(param.Method, param.Url, bytes.NewBuffer(body))
	if err != nil {

	}

	for key, value := range param.Headers {
		req.Header.Set(key, value.(string))
	}

	resp, err := client.Do(req)
	if err != nil {

	}

	return resp, nil
}

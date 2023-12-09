package logger

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type ILoggerService interface {
	InitLog(c *gin.Context) *Log
	HttpDefaultResponse(log *Log) gin.H
	HttpErrorResponse(log *Log, er error) gin.H
	WsDefaultResponse(log *Log) []byte
	WsErrorResponse(log *Log, err error) []byte
}

type loggerService struct {
	r iLoggerRepo
}

func NewLoggerService(r iLoggerRepo) *loggerService {
	err := r.MigrateLogger()
	if err != nil {
		return nil
	}

	return &loggerService{r: r}
}

func (s *loggerService) WsErrorResponse(log *Log, err error) []byte {
	log.Type = "ws"
	s.loggerRun(log, err.Error())
	return []byte(fmt.Sprintf(`{"status": "error", "error": "%s"}`, err.Error()))
}

func (s *loggerService) WsDefaultResponse(log *Log) []byte {
	log.Type = "ws"
	s.loggerRun(log, "")
	return []byte(`{"status": "success"}`)
}

func (s *loggerService) HttpErrorResponse(log *Log, er error) gin.H {
	log.Type = "http"
	s.loggerRun(log, er.Error())
	return gin.H{"status": "error", "error": er.Error()}
}

func (s *loggerService) HttpDefaultResponse(log *Log) gin.H {
	log.Type = "http"
	s.loggerRun(log, "")
	return gin.H{"status": "success"}
}

// /////////////////////////////////////////////////////////////////////////////////
func (s *loggerService) loggerRun(log *Log, errorText string) {
	log.Date_stop = time.Now()
	log.Milliseconds = int(log.Date_stop.Sub(log.Date_stop).Milliseconds())

	err := s.r.CreateLog(log)
	if err != nil {
		fmt.Println("----------------------------------------------------")
		fmt.Println(" -------------------- ERROR LOG --------------------")
		fmt.Println(err)
		fmt.Println("----------------------------------------------------")
	}

	if errorText != "" {
		err = s.r.CreateError(Error{Err_text: errorText, Date: time.Now()})
		if err != nil {
			fmt.Println("------------------------------------------------------")
			fmt.Println(" -------------------- ERROR ERROR --------------------")
			fmt.Println(err)
			fmt.Println("------------------------------------------------------")
		}
	}
}

func (s *loggerService) InitLog(c *gin.Context) *Log {
	log := Log{}

	log.Url = c.Request.Host + c.Request.URL.String()
	log.Method = c.Request.Method
	log.Date_start = time.Now()
	log.Ip = c.ClientIP()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Body_req = "ERROR PARSE BODY"
	}
	log.Body_req = string(body)

	for key, values := range c.Request.Header {
		log.Headers_resp = fmt.Sprintf("%s: %v", key, values)
	}

	return &log
}

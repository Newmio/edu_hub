package logger

import (
	"ed"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type ILoggerService interface {
	InitLog(c *gin.Context) *Log
	LoggerRun(log *Log, errorText string)
	HttpDefaultResponse(c *gin.Context, log *Log)
	HttpErrorResponse(c *gin.Context, log *Log, er error)
	HttpIdResponse(c *gin.Context, log *Log, id int)
	HttpBadAuthResponse(c *gin.Context, log *Log)
	HttpTokenResponse(c *gin.Context, log *Log, token, refresh string)
	HttpRegisterResponse(c *gin.Context, log *Log, token, refresh string, id_acc int)
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
	s.LoggerRun(log, err.Error())
	return []byte(fmt.Sprintf(`{"status": false, "error": "%s"}`, err.Error()))
}

func (s *loggerService) WsDefaultResponse(log *Log) []byte {
	log.Type = "ws"
	s.LoggerRun(log, "")
	return []byte(`{"status": true}`)
}

func (s *loggerService) HttpRegisterResponse(c *gin.Context, log *Log, token, refresh string, id_acc int){
	log.Type = "http"

	body, err := json.Marshal(gin.H{"status": true, "id": id_acc, "token": token, "refresh": refresh})
	if err != nil{
		log.Status = 500
		s.LoggerRun(log, ed.ErrTrace(err, ed.Trace()).Error())
	}
	log.Body_resp = string(body)
	log.Status = 200

	s.LoggerRun(log, "")

	c.Header("Request-Id", log.Request_id)
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"status": true, "id": id_acc, "token": token, "refresh": refresh})
}

func (s *loggerService) HttpTokenResponse(c *gin.Context, log *Log, token, refresh string){
	log.Type = "http"

	body, err := json.Marshal(gin.H{"status": true, "token": token, "refresh": refresh})
	if err != nil{
		log.Status = 500
		s.LoggerRun(log, ed.ErrTrace(err, ed.Trace()).Error())
	}
	log.Body_resp = string(body)
	log.Status = 200

	s.LoggerRun(log, "")

	c.Header("Request-Id", log.Request_id)
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"status": true, "token": token, "refresh": refresh})
}

func (s *loggerService) HttpBadAuthResponse(c *gin.Context, log *Log) {
	log.Type = "http"

	body, err := json.Marshal(gin.H{"status": false, "error": "unauthorized"})
	if err != nil{
		log.Status = 500
		s.LoggerRun(log, ed.ErrTrace(err, ed.Trace()).Error())
	}
	log.Body_resp = string(body)
	log.Status = 401

	s.LoggerRun(log, "unauthorized")

	c.Header("Request-Id", log.Request_id)
	c.Header("Content-Type", "application/json")
	c.AbortWithStatusJSON(401, gin.H{"status": false, "error": "unauthorized"})
}

func (s *loggerService) HttpIdResponse(c *gin.Context, log *Log, id int) {
	log.Type = "http"

	body, err := json.Marshal(gin.H{"status": true, "id": id})
	if err != nil{
		log.Status = 500
		s.LoggerRun(log, ed.ErrTrace(err, ed.Trace()).Error())
	}
	log.Body_resp = string(body)
	log.Status = 200

	s.LoggerRun(log, "")

	c.Header("Request-Id", log.Request_id)
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"status": true, "id": id})
}

func (s *loggerService) HttpErrorResponse(c *gin.Context, log *Log, er error) {
	log.Type = "http"
	log.Status = 500

	body, err := json.Marshal(gin.H{"status": false, "error": er.Error()})
	if err != nil{
		s.LoggerRun(log, ed.ErrTrace(err, ed.Trace()).Error())
	}
	log.Body_resp = string(body)

	s.LoggerRun(log, er.Error())

	c.Header("Request-Id", log.Request_id)
	c.Header("Content-Type", "application/json")
	c.AbortWithStatusJSON(500, gin.H{"status": false, "error": er.Error()})
}

func (s *loggerService) HttpDefaultResponse(c *gin.Context, log *Log) {
	log.Type = "http"

	body, err := json.Marshal(gin.H{"status": true})
	if err != nil{
		log.Status = 500
		s.LoggerRun(log, ed.ErrTrace(err, ed.Trace()).Error())
	}
	log.Body_resp = string(body)
	log.Status = 200

	s.LoggerRun(log, "")

	c.Header("Request-Id", log.Request_id)
	c.Header("Content-Type", "application/json")
	c.JSON(200, gin.H{"status": true})
}

// /////////////////////////////////////////////////////////////////////////////////
func (s *loggerService) LoggerRun(log *Log, errorText string) {
	log.Date_stop = time.Now()
	log.Milliseconds = int(log.Date_stop.Sub(log.Date_start).Milliseconds())

	if errorText == ""{
		log.Success = true
	}

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
	log.Request_id = fmt.Sprint(time.Now().UnixNano())

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

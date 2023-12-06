package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ILoggerService interface {
	DefaultResponse(log *Log)gin.H
	ErrorResponse(log *Log, er error)gin.H
	LoggerRun(log *Log, errorText string)
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

func (s *loggerService) ErrorResponse(log *Log, er error)gin.H{
	s.LoggerRun(log, er.Error())
	return gin.H{"status": "error", "error": er.Error()}
}

func (s *loggerService) DefaultResponse(log *Log)gin.H{
	s.LoggerRun(log, "")
	return gin.H{"status": "success"}
}

func (s *loggerService) LoggerRun(log *Log, errorText string) {
	log.Date_stop = time.Now()
	log.Milliseconds = uint(log.Date_stop.Sub(log.Date_stop).Milliseconds())

	err := s.r.CreateLog(log)
	if err != nil {
		fmt.Println("--------------------")
		fmt.Println(" ---- ERROR LOG ----")
		fmt.Println("--------------------")
	}

	if errorText != "" {
		err = s.r.CreateError(Error{Err_text: errorText, Date: time.Now()})
		if err != nil {
			fmt.Println("--------------------")
			fmt.Println(" --- ERROR ERROR ---")
			fmt.Println("--------------------")
		}
	}
}

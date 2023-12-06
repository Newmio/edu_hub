package logger

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

type Log struct {
	Id           uint      `json:"id" db:"id"`
	Url          string    `json:"url" db:"url"`
	Body_req     string    `json:"body_req" db:"body_req"`
	Headers_req  string    `json:"headers_req" db:"headers_req"`
	Status       uint      `json:"status" db:"status"`
	Body_resp    string    `json:"body_resp" db:"body_resp"`
	Headers_resp string    `json:"headers_resp" db:"headers_resp"`
	Method       string    `json:"method" db:"method"`
	Date_start   time.Time `json:"date_start" db:"date_start"`
	Date_stop    time.Time `json:"date_stop" db:"date_stop"`
	Milliseconds uint      `json:"milliseconds" db:"milliseconds"`
	Ip           string    `json:"ip" db:"ip"`
	Success      bool      `json:"success" db:"success"`
}

func (log *Log) CreateLog(c *gin.Context) {

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
}

type Error struct {
	Id       uint      `json:"id" db:"id"`
	Err_text string    `json:"error" db:"error"`
	Date     time.Time `json:"date" db:"date"`
}

package logger

import (
	"time"
)

type Log struct {
	Id           int       `json:"id" db:"id"`
	Url          string    `json:"url" db:"url"`
	Body_req     string    `json:"body_req" db:"body_req"`
	Headers_req  string    `json:"headers_req" db:"headers_req"`
	Status       int       `json:"status" db:"status"`
	Body_resp    string    `json:"body_resp" db:"body_resp"`
	Headers_resp string    `json:"headers_resp" db:"headers_resp"`
	Method       string    `json:"method" db:"method"`
	Date_start   time.Time `json:"date_start" db:"date_start"`
	Date_stop    time.Time `json:"date_stop" db:"date_stop"`
	Milliseconds int       `json:"milliseconds" db:"milliseconds"`
	Ip           string    `json:"ip" db:"ip"`
	Success      bool      `json:"success" db:"success"`
	Type string
}

type Error struct {
	Id       int       `json:"id" db:"id"`
	Err_text string    `json:"error" db:"error"`
	Date     time.Time `json:"date" db:"date"`
}

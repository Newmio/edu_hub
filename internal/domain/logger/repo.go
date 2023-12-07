package logger

import (
	"ed"
	"errors"

	"github.com/jmoiron/sqlx"
)

type iLoggerRepo interface {
	MigrateLogger() error
	CreateLog(log *Log) error
	CreateError(er Error) error
}

type loggerRepo struct {
	db *sqlx.DB
}

func NewLoggerRepo(db *sqlx.DB) *loggerRepo {
	return &loggerRepo{db: db}
}

func (db *loggerRepo) CreateLog(log *Log) error {
	str := `insert into logs(type, url, body_req, headers_req, status, body_resp, headers_resp,
		method, date_start, date_stop, milliseconds, ip, success) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`

	result, err := db.db.Exec(str, log.Type, log.Url, log.Body_req, log.Headers_req, log.Status, log.Body_resp,
		log.Headers_resp, log.Method, log.Date_start.Format("02-01-2006 15:04:05.000"),
		log.Date_stop.Format("02-01-2006 15:04:05.000"), log.Milliseconds, log.Ip, log.Success)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	if row == 0 {
		return errors.New("bad insert log")
	}

	return nil
}

func (db *loggerRepo) CreateError(er Error) error {
	str := `insert into errors(error, date) values($1,$2)`

	result, err := db.db.Exec(str, er.Err_text, er.Date.Format("02-01-2006 15:04:05.000"))
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	if row == 0 {
		return errors.New("bad insert log")
	}

	return nil
}

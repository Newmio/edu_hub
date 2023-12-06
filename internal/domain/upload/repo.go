package upload

import (
	"ed"
	"errors"

	"github.com/jmoiron/sqlx"
)

type iUploadRepo interface {
	MigrateFiles() error
	CreateFile(file FileHistory) error
}

type uploadRepo struct {
	db *sqlx.DB
}

func NewUploadRepo(db *sqlx.DB) *uploadRepo {
	return &uploadRepo{db: db}
}

func (db *uploadRepo) CreateFile(file FileHistory) error {
	str := "insert into files(id_account, directory, file, date, size) values($1,$2,$3,$4,$5)"

	result, err := db.db.Exec(str, file.Id_account, file.Directory, file.File, file.Date, file.Size)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	if row == 0 {
		return errors.New("bad insert file")
	}

	return nil
}

package lesson

import (
	"ed"
	"errors"

	"github.com/jmoiron/sqlx"
)

type iLessonRepo interface {
	MigrateLesson() error
	CreateLesson(lesson *Lesson) error
}

type lessonRepo struct {
	db *sqlx.DB
}

func NewLessonRepo(db *sqlx.DB) *lessonRepo {
	return &lessonRepo{db: db}
}

func (db *lessonRepo) CreateLesson(lesson *Lesson) error {
	str := `insert into lessons(name, date, id_teacher, description) values($1,$2,$3,$4)`

	result, err := db.db.Exec(str, lesson.Name, lesson.Date, lesson.Description)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	if row == 0 {
		return errors.New("bad insert lesson")
	}

	return nil
}

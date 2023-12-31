package lesson

import (
	"ed"
	"errors"

	"github.com/jmoiron/sqlx"
)

type iLessonRepo interface {
	MigrateLesson() error
	CreateLesson(lesson *Lesson) error
	CreateLessonVisit(visit *LessonVisit)error
	CreateSubject(sub *Subject)error
}

type lessonRepo struct {
	db *sqlx.DB
}

func NewLessonRepo(db *sqlx.DB) *lessonRepo {
	return &lessonRepo{db: db}
}

func (db *lessonRepo) CreateLesson(lesson *Lesson) error {
	str := `insert into lessons(name, date, id_teacher, id_subject, description) values($1,$2,$3,$4,$5)`

	result, err := db.db.Exec(str, lesson.Name, lesson.Date, lesson.Id_subject, lesson.Description)
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

func (db *lessonRepo) CreateSubject(sub *Subject)error{
	str := `insert into subjects(name, min_hours, max_hours, description) values($1,$2,$3,$4,$5)`

	result, err := db.db.Exec(str, sub.Name, sub.Min_hours, sub.Max_hours, sub.Description)
	if err != nil{
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil{
		return ed.ErrTrace(err, ed.Trace())
	}

	if row == 0{
		return errors.New("bad insert subject")
	}

	return nil
}

func (db *lessonRepo) CreateLessonVisit(visit *LessonVisit)error{
	str := `insert into lesson_visits(id_account, id_subject, id_lesson, hours, eval, behavior)
	values($1,$2,$3,$4,$5,$6)`

	result, err := db.db.Exec(str, visit.Id_acc, visit.Id_subject, visit.Id_lesson,
	visit.Hours, visit.Eval, visit.Behavior)
	if err != nil{
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil{
		return ed.ErrTrace(err, ed.Trace())
	}

	if row == 0{
		return errors.New("bad insert lesson_visit")
	}

	return nil
}

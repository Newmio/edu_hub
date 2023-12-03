package classroom

import (
	"ed"
	"errors"

	"github.com/jmoiron/sqlx"
)

type iClassroomRepo interface {
	MigrateClassroom() error
	CreateClassroom(class *Classroom) error
}

type classroomRepo struct {
	db *sqlx.DB
}

func NewClassroomRepo(db *sqlx.DB) *classroomRepo {
	return &classroomRepo{db: db}
}

func (db *classroomRepo) CreateClassroom(class *Classroom) error {
	str := `insert into classrooms(number, max_person, description, active) values($1,$2,$3,$4)`

	result, err := db.db.Exec(str, class.Number, class.Max_person, class.Description, class.Active)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	if row == 0 {
		return errors.New("bad insert classroom")
	}

	return nil
}

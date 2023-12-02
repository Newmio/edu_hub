package user

import (
	"ed"
	"errors"

	"github.com/jmoiron/sqlx"
)

type IUserRepo interface {
	CreatePerson(pers *Person) error
	CreateAccount(acc *Account) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (db *userRepo) CreatePerson(pers *Person) error {
	str := `insert into persons(name, last_name, middle_name, age, date, sex, role) values($1,$2,$3,$4,$5,$6,$7)`

	result, err := db.db.Exec(str, pers.Name, pers.Last_name, pers.Middle_name, pers.Age, pers.Date, pers.Sex, pers.Role)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	if row == 0 {
		return errors.New("bad insert person")
	}

	return nil
}

func (db *userRepo) CreateAccount(acc *Account) error {
	str := `insert into accounts(login, password, id_person, active) values($1,$2,$3,$4)`

	result, err := db.db.Exec(str, acc.Login, acc.Pass, acc.Id_person, acc.Active)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	if row == 0 {
		return errors.New("bad insert account")
	}

	return nil
}

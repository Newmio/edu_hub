package user

import (
	"ed"
	"errors"

	"github.com/jmoiron/sqlx"
)

type iUserRepo interface {
	MigrateUser() error
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
	tx, err := db.db.Beginx()
	if err != nil{
		return ed.ErrTrace(err, ed.Trace())
	}

	str := `select exists(select 1 from accounts where login = $1) as result`
	
	var res bool
	err = tx.QueryRow(str, acc.Login).Scan(&res)
	if err != nil{
		tx.Rollback()
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	if res{
		tx.Rollback()
		return errors.New("account already exists")
	}

	str = `insert into accounts(login, password, id_person, active) values($1,$2,$3,$4)`

	result, err := tx.Exec(str, acc.Login, acc.Pass, acc.Id_person, acc.Active)
	if err != nil {
		tx.Rollback()
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	if row == 0 {
		tx.Rollback()
		return errors.New("bad insert account")
	}

	if err := tx.Commit(); err != nil{
		tx.Rollback()
		return ed.ErrTrace(err, ed.Trace())
	}

	return nil
}

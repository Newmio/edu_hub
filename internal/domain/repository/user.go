package repository

import (
	"ed/internal/domain/model"
	"ed/pkg/util"
	"errors"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func newUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (db *userRepo) CreatePerson(pers *model.Person)error{
	str := `insert into persons(name, last_name, middle_name, age, date, sex, role) values($1,$2,$3,$4,$5,$6,$7)`

	result, err := db.db.Exec(str, pers.Name, pers.Last_name, pers.Middle_name, pers.Age, pers.Date, pers.Sex, pers.Role)
	if err != nil{
		return util.ErrDbTrace(err, str, util.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil{
		return util.ErrDbTrace(err, str, util.Trace())
	}

	if row == 0{
		return errors.New("bad insert person")
	}

	return nil
}

func (db *userRepo) CreateAccount(acc *model.Account) error {
	str := `insert into accounts(login, password, id_person, active) values($1,$2,$3,$4)`

	result, err := db.db.Exec(str, acc.Login, acc.Pass, acc.Id_person, acc.Active)
	if err != nil {
		return util.ErrDbTrace(err, str, util.Trace())
	}

	row, err := result.RowsAffected()
	if err != nil{
		return util.ErrDbTrace(err, str, util.Trace())
	}

	if row == 0{
		return errors.New("bad insert account")
	}

	return nil
}

package user

import (
	"ed"
	"errors"

	"github.com/jmoiron/sqlx"
)

type iUserRepo interface {
	MigrateUser() error
	CreatePerson(pers *Person, id_acc int) error
	CreateAccount(acc *Account) (int, error)
	GetAccountById(id int) (*Account, error)
	UpdateAccountAuthParams(id int, login, pass string) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (db *userRepo) UpdateAccountAuthParams(id int, login, pass string) error {
	if login != "" && pass != "" {

		str := "update accounts set login = $1, pass = $2 where id = $3"

		_, err := db.db.Exec(str, login, pass, id)
		if err != nil {
			return ed.ErrTrace(err, ed.Trace())
		}

	} else {

		if login != "" {
			str := "update accounts set login = $1 where id = $2"

			_, err := db.db.Exec(str, login, id)
			if err != nil {
				return ed.ErrTrace(err, ed.Trace())
			}

		}else if pass != ""{
			str := "update accounts set pass = $1 where id = $2"

			_, err := db.db.Exec(str, pass, id)
			if err != nil {
				return ed.ErrTrace(err, ed.Trace())
			}
		}else{
			return nil
		}
	}

	return nil
}

func (db *userRepo) GetAccountById(id int) (*Account, error) {
	var account Account

	str := "select * from accounts where id = $1"

	err := db.db.Get(&account, str, id)
	if err != nil {
		return nil, ed.ErrDbTrace(err, str, ed.Trace())
	}

	return &account, nil
}

func (db *userRepo) CreatePerson(pers *Person, id_acc int) error {
	tx, err := db.db.Beginx()
	if err != nil {
		return ed.ErrTrace(err, ed.Trace())
	}

	str := `insert into persons(name, last_name, middle_name, age, date, sex, role) values($1,$2,$3,$4,$5,$6,$7) returning id`

	var id_pers int

	err = tx.QueryRow(str, pers.Name, pers.Last_name, pers.Middle_name, pers.Age, pers.Date, pers.Sex, pers.Role).Scan(&id_pers)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	str = "update accounts set id_person = $1 where id = $2"

	_, err = tx.Exec(str, id_pers, id_acc)
	if err != nil {
		return ed.ErrDbTrace(err, str, ed.Trace())
	}

	return nil
}

func (db *userRepo) CreateAccount(acc *Account) (int, error) {
	tx, err := db.db.Beginx()
	if err != nil {
		return 0, ed.ErrTrace(err, ed.Trace())
	}

	str := `select exists(select 1 from accounts where login = $1) as result`

	var res bool
	err = tx.QueryRow(str, acc.Login).Scan(&res)
	if err != nil {
		tx.Rollback()
		return 0, ed.ErrDbTrace(err, str, ed.Trace())
	}

	if res {
		tx.Rollback()
		return 0, errors.New("account already exists")
	}

	str = `insert into accounts(login, password, id_person, role, active) values($1,$2,$3,$4,$5) returning id`

	var id int

	err = tx.QueryRow(str, acc.Login, acc.Pass, acc.Id_person, acc.Role, acc.Active).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, ed.ErrDbTrace(err, str, ed.Trace())
	}

	return id, nil
}

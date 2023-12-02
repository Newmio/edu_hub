package repository

import (
	"ed/internal/domain/model"

	"github.com/jmoiron/sqlx"
)

type Common interface {
	InitTables() error
	CheckExistInDb(table, param, value string) (bool, error)
}

type User interface {
	CreateAccount(acc *model.Account) error
	CreatePerson(pers *model.Person) error
}

type Repo struct {
	Common
	User
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		Common: newCommonRepo(db),
		User: newUserRepo(db),
	}
}

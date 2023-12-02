package service

import (
	"ed/internal/domain/model"
	"ed/internal/domain/repository"
)

type common interface {
	InitTables() error
	CheckExistInDb(table, param, value string) (bool, error)
}

type user interface {
	CreateAccount(acc *model.Account) error
	CreatePerson(pers *model.Person) error
}

type Service struct {
	common
	user
}

func NewService(r *repository.Repo) *Service {
	return &Service{
		common: r.Common,
		user: r.User,
	}
}

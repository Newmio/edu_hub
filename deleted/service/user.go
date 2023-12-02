package service

import (
	"ed/internal/domain/model"
	"ed/pkg/util"
	"errors"
)

func (s *Service) CreateAccount(acc *model.Account) error {

	result, err := s.common.CheckExistInDb("accounts", "login", acc.Login)
	if err != nil {
		return util.ErrTrace(err, util.Trace())
	}

	if !result {
		return s.user.CreateAccount(acc)
	}

	return errors.New("account already exists")
}

func (s *Service) CreatePerson(pers *model.Person) error {
	return s.user.CreatePerson(pers)
}

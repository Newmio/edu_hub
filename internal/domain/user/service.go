package user

type IUserService interface {
	CreateAccount(acc *Account) error
	CreatePerson(pers *Person) error
}

type userService struct {
	r IUserRepo
}

func NewUserService(r IUserRepo) *userService {
	return &userService{r: r}
}

func (s *userService) CreateAccount(acc *Account) error {

	// result, err := s.common.CheckExistInDb("accounts", "login", acc.Login)
	// if err != nil {
	// 	return util.ErrTrace(err, util.Trace())
	// }

	// if !result {
	// 	return s.user.CreateAccount(acc)
	// }

	// return errors.New("account already exists")

	return s.r.CreateAccount(acc)
}

func (s *userService) CreatePerson(pers *Person) error {
	return s.r.CreatePerson(pers)
}

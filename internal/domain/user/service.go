package user

type IUserService interface {
	CreateAccount(acc *Account) error
	CreatePerson(pers *Person) error
}

type userService struct {
	r iUserRepo
}

func NewUserService(r iUserRepo) *userService {
	err := r.MigrateUser()
	if err != nil{
		return nil
	}
	return &userService{r: r}
}

func (s *userService) CreateAccount(acc *Account) error {
	return s.r.CreateAccount(acc)
}

func (s *userService) CreatePerson(pers *Person) error {
	return s.r.CreatePerson(pers)
}

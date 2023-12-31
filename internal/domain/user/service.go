package user

import (
	"crypto/sha256"
	"ed"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IUserService interface {
	CreateAccount(acc *Account) (string, string, int, error)
	CreatePerson(pers *Person, id_acc int) error
	ParseToken(token string) (int, error)
}

type userService struct {
	r iUserRepo
}

func NewUserService(r iUserRepo) *userService {
	err := r.MigrateUser()
	if err != nil {
		return nil
	}
	return &userService{r: r}
}

func (s *userService) ParseToken(token string) (int, error) {
	t, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error signinig method")
		}

		return []byte(KEY), nil
	})
	if err != nil {
		return 0, ed.ErrTrace(err, ed.Trace())
	}

	claims, ok := t.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *userService) register(id int) (string, string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 800).Unix(),
		},
		id,
	})

	t, err := token.SignedString([]byte(KEY))
	if err != nil {
		return "", "", ed.ErrTrace(err, ed.Trace())
	}

	r, err := refresh.SignedString([]byte(KEY))
	if err != nil {
		return "", "", ed.ErrTrace(err, ed.Trace())
	}

	return t, r, nil
}

func (s *userService) CreateAccount(acc *Account) (string, string, int, error) {
	acc.Pass = generateHash(acc.Pass)

	if acc.Id_person < 1{
		acc.Active = false
	}

	if acc.Role == ""{
		acc.Role = "user"
	}

	id, err := s.r.CreateAccount(acc)
	if err != nil{
		return "", "", 0, ed.ErrTrace(err, ed.Trace())
	}

	token, refresh, err := s.register(id)
	if err != nil{
		return "", "", 0, ed.ErrTrace(err, ed.Trace())
	}

	return token, refresh, id, nil
}

func (s *userService) CreatePerson(pers *Person, id_acc int) error {

	t, err := time.Parse(ed.TIMEFORMAT, pers.Date)
	if err != nil{
		return ed.ErrTrace(err, ed.Trace())
	}

	pers.Age = time.Now().Year() - t.Year()
	
	return s.r.CreatePerson(pers, id_acc)
}

func generateHash(pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))

	return fmt.Sprint(hash.Sum([]byte(SALT)))
}

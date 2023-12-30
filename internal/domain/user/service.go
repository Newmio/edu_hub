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
	CreateAccount(acc *Account) (int, error)
	CreatePerson(pers *Person) error
	Register(login, pass string)(string, string, error)
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

func (s *userService) Register(login, pass string) (string, string, error) {

	account, err := s.r.GetAccount(login, generateHash(pass))
	if err != nil {
		return "", "", ed.ErrTrace(err, ed.Trace())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		account.Id,
	})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
		account.Id,
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

func (s *userService) CreateAccount(acc *Account) (int, error) {
	acc.Pass = generateHash(acc.Pass)
	return s.r.CreateAccount(acc)
}

func (s *userService) CreatePerson(pers *Person) error {
	return s.r.CreatePerson(pers)
}

func generateHash(pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))

	return fmt.Sprint(hash.Sum([]byte(SALT)))
}

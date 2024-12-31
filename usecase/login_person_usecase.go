package usecase

import (
	"api-mercearia-go/model"
	"api-mercearia-go/repository"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type LoginPersonUsecase struct {
	personRepository repository.PersonRepository
}

var (
	JWT_SECRET_KEY = "Y29mNjQ2ZDFhNTU0N2Y2OTQyNzZmMDU1NjM0NzljODIzNDg5ZTYxNmFlYzFkMjZkNjk3NmQ1Nw=="
)

func NewLoginPersonUsecase(personRepository repository.PersonRepository) LoginPersonUsecase {
	return LoginPersonUsecase{
		personRepository: personRepository,
	}
}

func (lu *LoginPersonUsecase) login(userLogin model.UserLogin) (string, error) {
	//logica

	return "", nil
}

func generateToken(person model.Person) (string, error) {

	clains := jwt.MapClaims{
		"id":       person.ID,
		"name":     person.Name,
		"userType": person.UserType,
		"email":    person.Email,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clains)
	tokenString, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	return tokenString, nil
}

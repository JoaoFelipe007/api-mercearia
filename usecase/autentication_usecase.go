package usecase

import (
	"api-mercearia-go/model"
	"api-mercearia-go/repository"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AutenticationUsecase struct {
	personRepository repository.PersonRepository
}

var (
	JWT_SECRET_KEY = "Y29mNjQ2ZDFhNTU0N2Y2OTQyNzZmMDU1NjM0NzljODIzNDg5ZTYxNmFlYzFkMjZkNjk3NmQ1Nw=="
)

func NewAutenticationUsecase(personRepository repository.PersonRepository) AutenticationUsecase {
	return AutenticationUsecase{
		personRepository: personRepository,
	}
}

func (lu *AutenticationUsecase) Login(userLogin model.UserLogin) (string, error) {

	person, err := lu.personRepository.GetPersonByEmail(userLogin.Username)
	if err != nil {
		return "", err
	}

	if person.ID != 0 {
		err := bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(userLogin.Password))
		if err != nil {
			return "", err
		}
		token, err := generateToken(person)

		if err != nil {
			return "", err
		}
		return token, nil
	}

	return "", nil

}

func generateToken(person model.Person) (string, error) {

	clains := jwt.MapClaims{
		"userID":   person.ID,
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

func VerifyToken(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// Verifique o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signature method")
		}

		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Valid Token")
}

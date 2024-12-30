package usecase

import (
	"api-mercearia-go/model"
	"api-mercearia-go/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PersonUsecase struct {
	personRepository repository.PersonRepository
}

func NewPersonUsecase(personRepository repository.PersonRepository) PersonUsecase {
	return PersonUsecase{
		personRepository: personRepository,
	}
}

func (pu *PersonUsecase) CreatePerson(person model.Person) (model.Person, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(person.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	person.Password = string(hash)

	personResult, err := pu.personRepository.CreatePerson(person)

	if err != nil {
		fmt.Print(err)
		return model.Person{}, err
	}

	return personResult, nil
}

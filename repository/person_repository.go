package repository

import (
	"api-mercearia-go/model"
	"database/sql"
	"log"
)

type PersonRepository struct {
	connection *sql.DB
}

func NewPersonRepository(connection *sql.DB) PersonRepository {
	return PersonRepository{
		connection: connection,
	}
}

func (pr *PersonRepository) CreatePerson(person model.Person) (model.Person, error) {
	var personResult model.Person

	query, err := pr.connection.Prepare(`
		INSERT INTO person (name, password, email, document, is_a_physical_person, user_type)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING *;
	`)

	if err != nil {
		log.Fatal(err)
		return model.Person{}, err
	}

	err = query.QueryRow(person.Name, person.Password, person.Email, person.Document, person.IsAPhysicalPerson, person.UserType).Scan(
		&personResult.ID,
		&personResult.Name,
		&personResult.Password,
		&personResult.Email,
		&personResult.Document,
		&personResult.IsAPhysicalPerson,
		&personResult.UserType,
		&personResult.RegistrationDate,
		&personResult.DateChange,
	)

	if err != nil {
		log.Fatal(err)
		return model.Person{}, err
	}

	query.Close()
	return personResult, nil
}

func (pr *PersonRepository) GetPersonByEmail(email string) (model.Person, error) {
	var personResult model.Person

	query := "SELECT * FROM person WHERE email = $1"

	err := pr.connection.QueryRow(query, email).Scan(
		&personResult.ID,
		&personResult.Name,
		&personResult.Password,
		&personResult.Email,
		&personResult.Document,
		&personResult.IsAPhysicalPerson,
		&personResult.UserType,
		&personResult.RegistrationDate,
		&personResult.DateChange,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return model.Person{}, nil
		}
		return model.Person{}, err
	}
	pr.connection.Close()
	return personResult, nil
}

package model

import "time"

type UserType int

const (
	Admin UserType = iota
	Manager
	Employee
)

type Person struct {
	ID                int        `json:"id"`
	Name              string     `json:"name"`
	Password          string     `json:"password"`
	Email             string     `json:"email"`
	Document          string     `json:"document"`
	IsAPhysicalPerson bool       `json:"isAPhysicalPerson"`
	UserType          UserType   `json:"userType"`
	RegistrationDate  time.Time  `json:"registrationDate"`
	DateChange        *time.Time `json:"dateChange"` // o * permite vim nulo
}

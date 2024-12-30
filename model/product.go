package model

import "time"

type Product struct {
	ID               int       `json:"id"`
	Description      string    `json:"description"`
	Priority         int       `json:"priority"`
	Price            float64   `json:"price"`
	Quantity         float64   `json:"quantity"`
	CategoryID       int       `json:"categoryId"`
	Status           bool      `json:"status"`
	dateChange       time.Time `json:"dateChange"`
	RegistrationDate time.Time `json:"registrationDate"`
}

package model

type Category struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Status      bool   `json:"status"`
}

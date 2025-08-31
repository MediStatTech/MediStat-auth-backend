package dto

import "time"

type Personal struct {
	PersonalID string    `json:"personal_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Phone      *string   `json:"phone"`
	Status     string    `json:"status"`
	Departure  string    `json:"departure"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AddPersonalRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Phone     *string `json:"phone"`
	Departure string  `json:"departure"`
	Password  string  `json:"password"`
}

type AddPersonalResponse struct{}

type GetPersonalRequest struct {
	PersonalID string `json:"personal_id"` // parse token todo
}

type GetPersonalResponse struct {
	Personal Personal `json:"personal"`
}

type ListPersonalRequest struct{}

type ListPersonalResponse struct {
	Personal []Personal `json:"personal"`
}

type UpdatePersonalStatusRequest struct {
	PersonalID string `json:"personal_id"`
	Status     string `json:"status"`
}

type UpdatePersonalStatusResponse struct {
	Personal Personal `json:"personal"`
}

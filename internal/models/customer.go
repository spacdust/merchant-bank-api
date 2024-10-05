package models

type Customer struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	HashedPassword string  `json:"hashed_password"`
	Balance        float64 `json:"balance"`
}

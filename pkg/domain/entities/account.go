package entities

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

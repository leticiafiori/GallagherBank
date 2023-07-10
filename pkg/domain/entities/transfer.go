package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transfer struct {
	Id                   uuid.UUID `json:"id"`
	IdAccountOrigin      uuid.UUID `json:"id_account_origin"`
	IdAccountDestination uuid.UUID `json:"id_account_destination"`
	Amount               int       `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

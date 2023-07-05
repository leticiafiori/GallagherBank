package account

import (
	"testing"

	"github.com/google/uuid"
)

type CreateAccountTestInput struct {
	ID        uuid.UUID
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt string
}

func TestCreateAccount(t *testing.T) {}

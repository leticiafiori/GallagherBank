package account

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
)

type CreateAccountTestInput struct {
	Id        uuid.UUID
	Name      string
	CPF       string
	Secret    string
	Balance   int
	CreatedAt time.Time
}

func TestCreateAccount(t *testing.T) {
	t.Run("Erro ao passar o balance zerado", func(t *testing.T) {

		esperado := errors.New("Balance cannot be last than zero")
		_, retorno := CreateAccount("Teste da Silva", "887.310.900-44", "123456", 0)

		if retorno == nil {
			t.Errorf("Expected error, but got nil")
		} else if retorno.Error() != esperado.Error() {
			t.Errorf("Expected error: %s, but got: %s", esperado, retorno)
		}

	})

	t.Run("Erro ao passar o balance negativo", func(t *testing.T) {

		esperado := errors.New("Balance cannot be last than zero")
		_, retorno := CreateAccount("Teste da Silva", "887.310.900-44", "123456", -2)

		if retorno == nil {
			t.Errorf("Expected error, but got nil")
		} else if retorno.Error() != esperado.Error() {
			t.Errorf("Expected error: %s, but got: %s", esperado, retorno)
		}

	})

	t.Run("Erro ao passar o CPF com formato inválido", func(t *testing.T) {

		esperado := errors.New("CPF format is incorrect")
		_, retorno := CreateAccount("Teste da Silva", "887.310.900-4433", "123456", 2)

		if retorno == nil {
			t.Errorf("Expected error, but got nil")
		} else if retorno.Error() != esperado.Error() {
			t.Errorf("Expected error: %s, but got: %s", esperado, retorno)
		}

	})

	t.Run("Criar a conta com CPF sem formatar ", func(t *testing.T) {
		newId, retorno := CreateAccount("Teste da Silva", "88712090044", "123456", 2)

		if retorno != nil {

			t.Errorf("Unexpected error: %v", retorno)
		}

		// Verificar se o ID retornado não é vazio
		if newId == uuid.Nil {
			t.Error("Expected non-empty UUID, but got nil")
		}

	})

	t.Run("Criar a conta com CPF formatado ", func(t *testing.T) {
		newId, retorno := CreateAccount("Teste da Silva", "887.120.900-44", "123456", 2)

		if retorno != nil {

			t.Errorf("Unexpected error: %v", retorno)
		}

		// Verificar se o ID retornado não é vazio
		if newId == uuid.Nil {
			t.Error("Expected non-empty UUID, but got nil")
		}

	})

}

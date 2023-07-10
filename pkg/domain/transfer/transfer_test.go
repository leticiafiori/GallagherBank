package transfer

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestCreateTransfer(t *testing.T) {
	t.Run("Erro ao passar o amount zerado", func(t *testing.T) {
		esperado := errors.New("Amount cannot be last than or equal to zero")
		retorno := CreateTransfer(uuid.MustParse("11047a8a-b24e-4de0-9521-5fbd28180b7a"), uuid.MustParse("80f656ca-a1f6-4e2c-a3ea-8ee447d97e58"), 0)

		if retorno == nil {
			t.Errorf("Expected error, but got nil")
		} else if retorno.Error() != esperado.Error() {
			t.Errorf("Expected error: %s, but got: %s", esperado, retorno)
		}
	})

}

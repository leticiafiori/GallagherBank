package account

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leticiafiori/GallegherBank/pkg/domain/entities"
)

func CreateAccount(name, cpf, secret string, balance int) (uuid.UUID, error) {
	// Gerar um ID único para a nova conta
	id, err := uuid.NewRandom()
	if err != nil {
		return Account{}, err
	}

	// Obter a data e hora atual
	createdAt := time.Now()
	// Criar uma nova instância da estrutura Account com os dados fornecidos
	account := Account{
		ID:        id,
		Name:      name,
		Cpf:       cpf,
		Secret:    secret,
		Balance:   balance,
		CreatedAt: createdAt,
	}

	// Realizar ações adicionais, como validação de dados ou interação com a camada de repositório, se necessário

	CheckBalance(account.Balance)

	// Retornar a nova conta criada
	return account, nil
}

func CheckBalance(amount int) error {
	if amount < 0 {
		return errors.New("Balance cannot be last than zero")
	}
	return nil
}

func GetAll() ([]entities.Account, error) {

}

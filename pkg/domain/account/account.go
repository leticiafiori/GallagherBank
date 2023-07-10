package account

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/leticiafiori/GallegherBank/pkg/domain/entities"
)

func CreateAccount(name, cpf, secret string, balance int) (uuid.UUID, error) {
	// Gerar um ID único para a nova conta
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.UUID{}, err
	}

	//Verificar se existe saldo para criar a conta
	err = CheckBalance(balance)
	if err != nil {
		return uuid.UUID{}, err
	}

	//Verificar se o CPF esta correto
	cpf, err = ValidaCPF(cpf)
	if err != nil {
		return uuid.UUID{}, err
	}

	// Obter a data e hora atual
	createdAt := time.Now()
	// Criar uma nova instância da estrutura Account com os dados fornecidos
	account := entities.Account{
		Id:        id,
		Name:      name,
		Cpf:       cpf,
		Secret:    secret,
		Balance:   balance,
		CreatedAt: createdAt,
	}

	// Retornar o id da  nova conta criada
	return account.Id, nil
}

func CheckBalance(balance int) error {
	if balance <= 0 {
		return errors.New("Balance cannot be last than zero")
	}
	return nil
}

func ValidaCPF(cpf string) (string, error) {
	if len(cpf) == 11 || len(cpf) == 14 {
		if len(cpf) == 11 {
			return cpf, nil
		} else {

			if string([]rune(cpf)[3]) == "." && string([]rune(cpf)[7]) == "." && string([]rune(cpf)[11]) == "-" {
				cpf = strings.Replace(cpf, ".", "", 2)
				cpf = strings.Replace(cpf, "-", "", 1)
				return cpf, nil
			}

		}
	} else {
		return cpf, errors.New("CPF format is incorrect")
	}
	return cpf, nil
}

func GetAll() ([]entities.Account, error) {
	return []entities.Account{}, nil
}

func GetBalance(cpf string) (int, error) {
	return 0, nil
}

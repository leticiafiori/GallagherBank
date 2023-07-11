package transfer

import (
	"errors"

	"github.com/google/uuid"

	"github.com/leticiafiori/GallegherBank/pkg/domain/entities"
)

func CreateTransfer(idOrigin, idDestination uuid.UUID, amount int) error {

	//busca a conta de origin
	accountOrigin, err := GetAccount(idOrigin)
	if err != nil {
		return err
	}

	//verifica se o idDestiantion existe
	err = CheckAccountExist(idDestination)
	if err != nil {
		return err
	}

	//busca a conta de destino
	accountDestination, err := GetAccount(idDestination)
	if err != nil {
		return err
	}

	//verifica se o amount(valor de transferencia) é maior que zero
	err = CheckAmount(amount)
	if err != nil {
		return err
	}

	//verifica se as contas são diferentes
	err = CheckDifferentId(idOrigin, idDestination)
	if err != nil {
		return err
	}

	//verificar se o balance da conta de origem é maior que o ammount

	//atualiza o saldo as duas contas
	UpdateBalance(accountOrigin, accountDestination, amount)

	return nil

}

func UpdateBalance(accountOrigin, accouuntDestination entities.Account, amount int) {
	accountOrigin.Balance = accountOrigin.Balance - amount
	accouuntDestination.Balance = accouuntDestination.Balance + amount

}

func CheckAmount(amount int) error {
	if amount <= 0 {

		return errors.New("Amount cannot be last than or equal to zero")
	}
	return nil
}

func CheckAccountExist(idAccount uuid.UUID) error {

	//buscar no banco se existe essa conta existe

	//	return errors.New("Account not found")

	return nil
}

func CheckDifferentId(idOrigin, idDestination uuid.UUID) error {
	if idOrigin == idDestination {
		return errors.New("Account Origin is equal to account Destination")
	}
	return nil
}

func GetAccount(id uuid.UUID) (entities.Account, error) {
	//busca a conta

	return entities.Account{}, nil
}

func GetAllTransfer(idOrigin uuid.UUID) ([]entities.Transfer, error) {
	return []entities.Transfer{}, nil

}

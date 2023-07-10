package transfer

import (
	"github.com/google/uuid"
	"github.com/leticiafiori/GallegherBank/pkg/domain/entities"
)

func CreateTransfer(idOrigin, idDestination uuid.UUID, amount int) error {

	//verifica se o idDestiantion existe

	//verifica se o amount(valor de transferencia) é maior que zero

	//verificar se o balance da conta de origem é maior que o ammount

	//atualiza o saldo as duas contas

	return nil

}

func GetAllTransfer(idOrigin uuid.UUID) ([]entities.Transfer, error) {
	return []entities.Transfer{}, nil

}

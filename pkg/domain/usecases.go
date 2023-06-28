package domain

import (
	"github.com/google/uuid"
	"github.com/leticiafiori/GallegherBank/pkg/domain/entities"
)

type AccountUseCase interface {
	GetAll() ([]entities.Account, error)
	CreateAccount(name string, cpf string, secret string, balance int) (uuid.UUID, error)
	GetBalance(cpf string) (int, error)
}

type LoginUseCase interface {
	CheckLogin(cpf string, secret string) (uuid.UUID, error)
}

type TranferUserCase interface {
	GetAllTranfer(idAccountOrigin uuid.UUID) ([]entities.Tranfer, error)
	CreateTransfer(IdAccountOrigin uuid.UUID, IdAccountDestination uuid.UUID, amount int) error
}

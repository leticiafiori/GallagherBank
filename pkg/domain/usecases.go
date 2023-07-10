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

type TransferUserCase interface {
	GetAllTransfer(idAccountOrigin uuid.UUID) ([]entities.Transfer, error)
	CreateTransfer(IdAccountOrigin uuid.UUID, IdAccountDestination uuid.UUID, amount int) error
}

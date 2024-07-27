package service

import (
	"context"
	api "github.com/alserov/smart_contract/internal/contracts"
	"github.com/alserov/smart_contract/internal/service/models"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Service interface {
	Contract
}

type Contract interface {
	Deposit(ctx context.Context, dep models.Deposit) error
	GetBalance(ctx context.Context) (float64, error)
	Withdraw(ctx context.Context, wth models.Withdraw) error
}

type ContractParams struct {
	Conn *api.Api
	Cl   *ethclient.Client
}

func NewService(contr ContractParams) Service {
	return &service{
		contract{
			cl:   contr.Cl,
			conn: contr.Conn,
		},
	}
}

type service struct {
	contract
}

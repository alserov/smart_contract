package service

import (
	"context"
	"fmt"
	api "github.com/alserov/smart_contract/internal/contracts"
	"github.com/alserov/smart_contract/internal/service/models"
	"github.com/alserov/smart_contract/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type contract struct {
	cl   *ethclient.Client
	conn *api.Api
}

func (c contract) Deposit(ctx context.Context, dep models.Deposit) error {
	auth, err := api.GetAccountAuth(ctx, c.cl, dep.From)
	if err != nil {
		return fmt.Errorf("failed to get account auth: %w", err)
	}

	opts, err := c.conn.Deposit(auth, big.NewInt(dep.Amount))
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	if err = api.CheckTransactionReceipt(c.cl, opts.Hash().String()); err != nil {
		return fmt.Errorf("failed to deposit: %w", err)
	}

	return nil
}

func (c contract) GetBalance(ctx context.Context) (float64, error) {
	reply, err := c.conn.Balance(&bind.CallOpts{})
	if err != nil {
		return 0, utils.NewError(err.Error(), utils.Internal)
	}

	value, _ := reply.Float64()

	return value, nil
}

func (c contract) Withdraw(ctx context.Context, wth models.Withdraw) error {
	auth, err := api.GetAccountAuth(ctx, c.cl, wth.To)
	if err != nil {
		return fmt.Errorf("failed to get account auth: %w", err)
	}

	opts, err := c.conn.Withdraw(auth, big.NewInt(wth.Amount))
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	if err = api.CheckTransactionReceipt(c.cl, opts.Hash().String()); err != nil {
		return fmt.Errorf("failed to withdraw: %w", err)
	}

	return nil
}

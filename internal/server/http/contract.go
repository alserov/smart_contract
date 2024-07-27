package http

import (
	"encoding/json"
	"fmt"
	"github.com/alserov/smart_contract/internal/service"
	"github.com/alserov/smart_contract/internal/service/models"
	"github.com/alserov/smart_contract/internal/utils"
	"net/http"
)

type Contract struct {
	srvc service.Service
}

func NewContractHandler(srvc service.Contract) Contract {
	return Contract{
		srvc: srvc,
	}
}

// GetBalance godoc
// @Summary      Get wallet balance
// @Description  Get wallet balance
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Success      200  {object} int
// @Failure      400  {object} string
// @Failure      500  {object} string
// @Router       /balance [get]
func (c *Contract) GetBalance(w http.ResponseWriter, r *http.Request) error {
	reply, err := c.srvc.GetBalance(r.Context())
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}

	b, err := json.Marshal(reply)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)

	return nil
}

// Deposit godoc
// @Summary      Deposit wallet balance
// @Description  Deposit wallet balance
// @Param 		 input body models.Deposit true "deposit"
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Success      200  {object} string
// @Failure      400  {object} string
// @Failure      500  {object} string
// @Router       /deposit [post]
func (c *Contract) Deposit(w http.ResponseWriter, r *http.Request) error {
	var dep models.Deposit
	if err := json.NewDecoder(r.Body).Decode(&dep); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	err := c.srvc.Deposit(r.Context(), dep)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

// Withdraw godoc
// @Summary      Withdraw wallet balance
// @Description  Withdraw wallet balance
// @Param 		 input body models.Withdraw true "withdraw"
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Success      200  {object} string
// @Failure      400  {object} string
// @Failure      500  {object} string
// @Router       /withdraw [post]
func (c *Contract) Withdraw(w http.ResponseWriter, r *http.Request) error {
	var wth models.Withdraw
	if err := json.NewDecoder(r.Body).Decode(&wth); err != nil {
		return utils.NewError(err.Error(), utils.BadRequest)
	}

	err := c.srvc.Withdraw(r.Context(), wth)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

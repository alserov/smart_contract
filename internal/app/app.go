package app

import (
	"context"
	"fmt"
	"github.com/alserov/smart_contract/internal/config"
	api "github.com/alserov/smart_contract/internal/contracts"
	"github.com/alserov/smart_contract/internal/logger"
	ctrl "github.com/alserov/smart_contract/internal/server/http"
	"github.com/alserov/smart_contract/internal/service"
	"net/http"
	"os/signal"
	"syscall"

	_ "github.com/alserov/smart_contract/docs"
)

const (
	port = 5000
)

func MustStart(cfg *config.Config) {
	conn, cl := api.MustSetupContract(cfg.ContractAddr)

	srvc := service.NewService(service.ContractParams{
		Conn: conn,
		Cl:   cl,
	})

	m := http.NewServeMux()

	l := logger.NewSlog()

	ctrl.SetupRoutes(m, l, ctrl.Controller{
		Contract: ctrl.NewContractHandler(srvc),
	})

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		l.Info(fmt.Sprintf("starting server on: %d", port))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), m); err != nil {
			panic("failed to start server: " + err.Error())
		}
	}()

	<-ctx.Done()
	l.Info("server shutdown")
}

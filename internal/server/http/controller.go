package http

import (
	"github.com/alserov/smart_contract/internal/logger"
	"github.com/alserov/smart_contract/internal/middleware"
	"net/http"

	"github.com/swaggo/http-swagger/v2"
)

type Controller struct {
	Contract
}

func SetupRoutes(m *http.ServeMux, l logger.Logger, ctrl Controller) {
	m.Handle("GET /swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:1323/swagger/doc.json")))

	m.HandleFunc("GET /v1/balance", middleware.WithRecovery(middleware.WithErrorHandler(ctrl.GetBalance)))
	m.HandleFunc("POST /v1/withdraw", middleware.WithRecovery(middleware.WithErrorHandler(ctrl.Withdraw)))
	m.HandleFunc("POST /v1/deposit", middleware.WithRecovery(middleware.WithErrorHandler(ctrl.Deposit)))
}

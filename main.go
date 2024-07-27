package main

import (
	"github.com/alserov/smart_contract/internal/app"
	"github.com/alserov/smart_contract/internal/config"
)

// @title Wallet api
// @version 1.0
// @description Wallet that interacts with smart contract
// @BasePath /v1
func main() {
	app.MustStart(config.MustLoad())
}

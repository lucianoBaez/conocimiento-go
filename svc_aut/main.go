package main

import (
	"github.com/d-Una-Interviews/svc_aut/pkg/api"
	_ "github.com/dimiro1/banner/autoload"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
)

func main() {

	// Create a new Server
	api.NewServer()

	// Initialize Web REST API
	api.InitAPI()

	// Start the Server
	api.StartServer()
}

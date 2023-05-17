package main

import (
	"github.com/leandronowras/go-oportunities/config"
	"github.com/leandronowras/go-oportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}

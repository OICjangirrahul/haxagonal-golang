package main

import (
	"log"

	v1 "github.com/OICjangirrahul/haxagonal-golang/cmd/v1"
	"github.com/OICjangirrahul/haxagonal-golang/config"

	"github.com/gin-gonic/gin"

	"github.com/OICjangirrahul/haxagonal-golang/internal"
)

func main() {
	r := gin.Default()

	config.ReadConfig()

	app, err := internal.InitializeApp()
	if err != nil {
		log.Fatalf("failed to initialize controller: %v", err)
	}

	// Setup routes
	v1.SetupRoutes(r, app.Controllers)

	// Start the server
	if err := r.Run(":" + config.Cfg.Server.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

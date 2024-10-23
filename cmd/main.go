package main

import (
	"log"

	v1 "github.com/OICjangirrahul/haxagonal-golang/cmd/v1"
	"github.com/OICjangirrahul/haxagonal-golang/config"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/OICjangirrahul/haxagonal-golang/docs"
	"github.com/OICjangirrahul/haxagonal-golang/internal"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	config.ReadConfig()

	app, err := internal.InitializeApp()
	if err != nil {
		log.Fatalf("failed to initialize controller: %v", err)
	}

	// Setup routes
	v1.SetupRoutes(r, app.Controllers)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Start the server
	if err := r.Run(":" + config.Cfg.Server.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

package v1

import (
	"github.com/OICjangirrahul/haxagonal-golang/internal/adapters/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, c *controller.Controllers) {
	v1 := r.Group("/api/v1/")
	{
		v1.POST("/accounts", c.AccountController.CreateAccount)

		v1.POST("/auth/saveToken", c.AuthController.SaveToken)
	}
}

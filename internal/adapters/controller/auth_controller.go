package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/in"
	"github.com/OICjangirrahul/haxagonal-golang/internal/validation"
)

type authRequest struct {
	Token string `json:"token" binding:"required,min=8"`
}

type AuthController struct {
	AuthInPort in.AuthInPort
	Validator  *validator.Validate
}

func NewAuthController(authInPort in.AuthInPort, validator *validator.Validate) *AuthController {
	return &AuthController{
		AuthInPort: authInPort,
		Validator:  validator,
	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping account
// @Schemes
// @Description do ping
// @Tags portal-api
// @Accept json
// @Produce json
// @Success 200 {string} Token created successfully
// @Router /auth/saveToken [post]
// @Param saveToken body authRequest true "saveToken Info"
func (c *AuthController) SaveToken(ctx *gin.Context) {
	var authRequest authRequest

	if err := validation.ValidateRequest(ctx, c.Validator, &authRequest); err != nil {
		return
	}

	err := c.AuthInPort.SaveToken(authRequest.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Token created successfully"})

}

package controller

import (
	"net/http"

	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/in"
	"github.com/OICjangirrahul/haxagonal-golang/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	AuthInPort in.AuthInPort
	Validator  *validator.Validate
}

func NewAuthController(authInPort in.AuthInPort, validator *validator.Validate) *AuthController {
	return &AuthController{
		AuthInPort: authInPort, // Use the correct field name with uppercase "A"
		Validator:  validator,  // Correctly assign validator
	}
}

func (c *AuthController) SaveToken(ctx *gin.Context) {
	var authRequest struct {
		Token string `json:"token" binding:"required,min=8"`
	}

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

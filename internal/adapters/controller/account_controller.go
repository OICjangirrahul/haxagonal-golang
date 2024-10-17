package controller

import (
	"net/http"

	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/in"
	"github.com/OICjangirrahul/haxagonal-golang/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AccountController struct {
	AccountInPort in.AccountInPort
	Validator     *validator.Validate
}

func NewAccountController(accountInPort in.AccountInPort, validator *validator.Validate) *AccountController {
	return &AccountController{
		AccountInPort: accountInPort,
		Validator:     validator,
	}
}

func (c *AccountController) CreateAccount(ctx *gin.Context) {

	var accountRequest struct {
		Username string `json:"username" binding:"required,min=3"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := validation.ValidateRequest(ctx, c.Validator, &accountRequest); err != nil {
		return
	}

	err := c.AccountInPort.CreateAccount(accountRequest.Username, accountRequest.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Account created successfully"})
}

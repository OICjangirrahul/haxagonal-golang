package controller

type Controllers struct {
	AccountController *AccountController
	AuthController    *AuthController
}

func NewControllers(accountController *AccountController, authController *AuthController) *Controllers {
	return &Controllers{
		AccountController: accountController,
		AuthController:    authController,
	}
}

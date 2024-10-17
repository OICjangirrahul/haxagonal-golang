package usecase

import (
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/domain"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/in"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/out"
)

type AuthUseCase struct {
	authRepo out.AuthOutPort // Out port interface
}

func NewAuthUseCase(authRepo out.AuthOutPort) in.AuthInPort {
	return &AuthUseCase{authRepo: authRepo}
}

// CreateAccount implements the account creation logic
func (uc *AuthUseCase) SaveToken(token string) error { // Implement SaveToken method
	redisToken := &domain.AuthToken{Token: token}
	return uc.authRepo.TokenSave(redisToken) // Use the authRepo to save the token
}

// internal/application/service/auth_service.go

package service

import (
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/domain"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/in"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/out"
)

type AuthRegistrationService struct {
	authOutPort out.AuthOutPort // Out port interface
}

// NewAuthRegistrationService initializes a new AuthRegistrationService
func NewAuthRegistrationService(authOutPort out.AuthOutPort) in.AuthInPort {
	return &AuthRegistrationService{authOutPort: authOutPort}
}

// SaveToken implements the token saving logic
func (s *AuthRegistrationService) SaveToken(token string) error {
	redisToken := &domain.AuthToken{Token: token}
	return s.authOutPort.TokenSave(redisToken) // Use authOutPort to save the token
}

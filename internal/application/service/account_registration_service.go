package service

import (
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/domain"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/in"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/out"
)

type AccountRegistrationService struct {
	accountOutPort out.AccountOutPort // Out port interface
}

// NewAccountRegistrationService initializes a new AccountRegistrationService
func NewAccountRegistrationService(accountOutPort out.AccountOutPort) in.AccountInPort {
	return &AccountRegistrationService{accountOutPort: accountOutPort}
}

// CreateAccount implements the account registration logic
func (s *AccountRegistrationService) CreateAccount(username, password string) error {
	account := &domain.Account{Username: username, Password: password}
	return s.accountOutPort.Save(account)
}

package usecase

import (
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/domain"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/in"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/out"
)

type AccountUseCase struct {
	accountRepo out.AccountOutPort // Out port interface
}

// NewAccountUseCase initializes a new AccountUseCase
func NewAccountUseCase(accountRepo out.AccountOutPort) in.AccountInPort {
	return &AccountUseCase{accountRepo: accountRepo}
}

// CreateAccount implements the account creation logic
func (uc *AccountUseCase) CreateAccount(username, password string) error {
	account := &domain.Account{Username: username, Password: password}
	return uc.accountRepo.Save(account)
}

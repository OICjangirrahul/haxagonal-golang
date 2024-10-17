package repository

import (
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/domain"
	"github.com/OICjangirrahul/haxagonal-golang/internal/application/port/out"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

// NewAccountRepository initializes a new AccountRepository
func NewAccountRepository(db *gorm.DB) out.AccountOutPort {
	return &AccountRepository{db: db}
}

// Save saves a new account to the database
func (r *AccountRepository) Save(account *domain.Account) error {
	return r.db.Create(account).Error
}

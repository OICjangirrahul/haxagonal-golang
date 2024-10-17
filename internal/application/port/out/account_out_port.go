package out

import "github.com/OICjangirrahul/haxagonal-golang/internal/application/domain"

type AccountOutPort interface {
	Save(account *domain.Account) error
}

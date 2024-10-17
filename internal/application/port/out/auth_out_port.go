// internal/application/port/out/auth_out_port.go

package out

import "github.com/OICjangirrahul/haxagonal-golang/internal/application/domain"

type AuthOutPort interface {
	TokenSave(token *domain.AuthToken) error
}

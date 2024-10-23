package in

type AuthInPort interface {
	SaveToken(token string) error
}

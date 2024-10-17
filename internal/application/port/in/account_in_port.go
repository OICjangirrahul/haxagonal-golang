package in

type AccountInPort interface {
    CreateAccount(username, password string) error
}

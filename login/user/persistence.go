package user

type AuthenticationVerifier interface {
	CheckUser(Username, Password) (bool, error)
}

type User struct {
	Username
	Password
}

type Username string
type Password string

// user/repository.go
package user

type User struct {
	ID    string
	Email string
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
}

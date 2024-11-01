package user

type UserProvider interface {
	Create(user User) (int, error)
}

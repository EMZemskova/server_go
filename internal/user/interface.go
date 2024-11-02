package user

type Provider interface {
	Create(user User) (int, error)
}

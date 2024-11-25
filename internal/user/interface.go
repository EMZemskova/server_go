package user

type Provider interface {
	Create(user User) (int, error)

	GetStat(id int64) (Statistics, error)

	GetStats() (map[int64]Statistics, error)
}

package stats

type Provider interface {
	GetStat(id int64) (Statistics, error)

	GetStats() ([]Statistics, error)
}

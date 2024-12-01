package stats

type Provider interface {
	GetStat(id int64) (Statistics, error)

	GetStats() ([]Statistics, error)
}

type Cacher interface {
	Provider
	CacheStat(id int64) (Statistics, error)

	CacheStats() (map[int64]Statistics, error)
}

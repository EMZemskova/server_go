package stats

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

type cacher struct {
	Provider
	mu             sync.RWMutex
	lastUpdate     time.Time
	userStatistics map[int64]Statistics
}

func NewCache(provider Provider) *cacher {
	return &cacher{
		Provider:       provider,
		userStatistics: make(map[int64]Statistics),
	}
}

func (c *cacher) StartCacheUpdater() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		<-ticker.C
		_, err := c.GetStats()
		if err != nil {
			errors.Wrap(err, "failed to get stat after 30s")
		}
	}
}

func (c *cacher) GetStat(id int64) (Statistics, error) {
	userStat, err := c.Provider.GetStat(id)
	if err != nil {
		return Statistics{}, errors.Wrap(err, "failed to get stats for person")
	}
	return userStat, err
}

func (c *cacher) GetStats() ([]Statistics, error) {
	c.mu.RLock()
	if time.Since(c.lastUpdate) < 30*time.Second {
		c.mu.RUnlock()
		return c.mapToSlice(c.userStatistics), nil
	}
	c.mu.RUnlock()
	stats, err := c.Provider.GetStats()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stats for caching")
	}
	c.SetStats(stats)
	return c.mapToSlice(c.userStatistics), err
}

func (c *cacher) mapToSlice(statsMap map[int64]Statistics) []Statistics {
	c.mu.RLock()
	defer c.mu.RUnlock()
	statsSlice := make([]Statistics, 0, len(statsMap))
	for _, stat := range statsMap {
		statsSlice = append(statsSlice, stat)
	}
	return statsSlice
}

func (c *cacher) SetStats(stats []Statistics) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, stat := range stats {
		c.userStatistics[stat.ID] = stat
	}
}

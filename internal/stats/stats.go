package stats

import (
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type stats struct {
	db *gorm.DB
}

func New(db *gorm.DB) *stats {
	return &stats{db: db}
}

func (s *stats) StartCacheUpdater() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		_, err := s.CacheStats()
		if err != nil {
			errors.Wrap(err, "failed to get stat after 30s")
		}
	}
}

func (s *stats) GetStat(id int64) (Statistics, error) {
	var stats Statistics
	query :=
		`SELECT u.id, u.username, COUNT(m.id) FILTER (WHERE m.sender = u.id) AS write_message, COUNT(DISTINCT c.id) FILTER (WHERE c.creator = u.id OR c.guest = u.id) AS chats_in    
		 FROM users u    
		 LEFT JOIN messages m ON m.sender = u.id    
		 LEFT JOIN chats c ON c.creator = u.id OR c.guest = u.id    
		 WHERE u.id = ?  
		 GROUP BY u.id, u.username;`
	if err := s.db.Raw(query, id).Scan(&stats).Error; err != nil {
		return Statistics{}, errors.Wrap(err, "failed to create user")
	}
	return stats, nil
}

func (s *stats) CacheStat(id int64) (Statistics, error) {
	userStat, err := s.GetStat(id)
	if err != nil {
		return Statistics{}, errors.Wrap(err, "failed to get stats for person")
	}
	return userStat, err
}

func (s *stats) GetStats() ([]Statistics, error) {
	var stats []Statistics
	query :=
		`SELECT u.id, u.username, COUNT(m.id) FILTER (WHERE m.sender = u.id) AS write_message, COUNT(DISTINCT c.id) FILTER (WHERE c.creator = u.id OR c.guest = u.id) AS chats_in
		 FROM users u
         LEFT JOIN messages m ON m.sender = u.id
         LEFT JOIN chats c ON c.creator = u.id OR c.guest = u.id
         GROUP BY u.id, u.username;`
	if err := s.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get people statistics")
	}
	return stats, nil
}

func (s *stats) CacheStats() (map[int64]Statistics, error) {
	MU.RLock()
	if time.Since(lastUpdate) < 30*time.Second {
		MU.RUnlock()
		return UserStatistics, nil
	}
	MU.RUnlock()
	stats, err := s.GetStats()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stats for caching")
	}
	MU.Lock()
	defer MU.Unlock()
	for _, stat := range stats {
		UserStatistics[stat.ID] = stat
	}
	return UserStatistics, err
}

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

func (s *stats) startCacheUpdater() {
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
		`SELECT u.id, u.username, COALESCE(message_counts.write_message, 0) AS write_message, COALESCE(chat_counts.chats_in, 0) AS chats_in  
		FROM users u 
				LEFT JOIN (
    				SELECT m.sender, COUNT(*) AS write_message 
    				FROM messages m 
    				GROUP BY m.sender
				) AS message_counts ON message_counts.sender = u.id 
				LEFT JOIN (
    				SELECT c.creator AS user_id, COUNT(DISTINCT c.id) AS chats_in 
    				FROM chats c 
    				GROUP BY c.creator
    				UNION ALL
    					SELECT c.guest AS user_id, COUNT(DISTINCT c.id) AS chats_in 
    					FROM chats c 
    					GROUP BY c.guest
				) AS chat_counts ON chat_counts.user_id = u.id  
		WHERE u.id = ?  
		GROUP BY u.id, u.username, message_counts.write_message, chat_counts.chats_in;`
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
		`SELECT u.id, u.username, COALESCE(message_counts.write_message, 0) AS write_message, COALESCE(chat_counts.chats_in, 0) AS chats_in    
		FROM users u   
			LEFT JOIN (  
    			SELECT  m.sender, COUNT(*) AS write_message   
    			FROM messages m   
    			GROUP BY m.sender  
			) AS message_counts ON message_counts.sender = u.id   
			LEFT JOIN (  
    			SELECT user_id, COUNT(DISTINCT id) AS chats_in   
   				FROM (
        			SELECT c.creator AS user_id, c.id
        			FROM chats c
        			UNION ALL  
        			SELECT c.guest AS user_id, c.id
        			FROM chats c  
    		) AS combined_chats
    		GROUP BY user_id  
			) AS chat_counts ON chat_counts.user_id = u.id    
		GROUP BY u.id, u.username, message_counts.write_message, chat_counts.chats_in;`
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

package stats

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type stats struct {
	db *gorm.DB
}

func NewProvider(db *gorm.DB) *stats {
	return &stats{db: db}
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

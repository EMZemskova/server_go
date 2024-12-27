package stats

import (
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type stats struct {
	db *pgx.Conn
}

func NewProvider(db *pgx.Conn) *stats {
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
	if err := s.db.QueryRow(query, id).Scan(
		&stats.ID,
		&stats.Username,
		&stats.WriteMessage,
		&stats.ChatsIn,
	); err != nil {
		logrus.Error("failed to get statistics", err)
		return Statistics{}, errors.Wrap(err, "failed to get statistics")
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
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query for statistics")
	}
	defer rows.Close()
	for rows.Next() {
		var stat Statistics
		if err := rows.Scan(&stat.ID, &stat.Username, &stat.WriteMessage, &stat.ChatsIn); err != nil {
			logrus.Error("failed to scan row into statistics struct", err)
			return nil, errors.Wrap(err, "failed to scan row into statistics struct")
		}
		stats = append(stats, stat)
	}
	if rows.Err() != nil {
		logrus.Error("error iterating over rows", err)
		return nil, errors.Wrap(rows.Err(), "error iterating over rows")
	}
	return stats, nil
}

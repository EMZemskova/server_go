package user

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func New(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) Create(user User) (int, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to create user")
	}
	return int(user.ID), nil
}

func (u *user) GetStat(id int64) (Statistics, error) {
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
	if err := u.db.Raw(query, id).Scan(&stats).Error; err != nil {
		return Statistics{}, errors.Wrap(err, "failed to create user")
	}
	return stats, nil
}

func (u *user) GetStats() (map[int64]Statistics, error) {
	var stats []Statistics
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
        GROUP BY u.id, u.username, message_counts.write_message, chat_counts.chats_in;`
	if err := u.db.Raw(query).Scan(&stats).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get people statistics")
	}
	mu.Lock()
	defer mu.Unlock()
	for _, stat := range stats {
		userStatistics[stat.ID] = stat
	}
	return userStatistics, nil
}

package message

import "time"

type Message struct {
	ID        int64
	Chat      int64
	Sender    int64
	Text      string
	DeletedAt *time.Time
}

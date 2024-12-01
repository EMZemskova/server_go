package stats

import (
	"sync"
	"time"
)

type Statistics struct {
	ID           int64
	Username     string
	ChatsIn      int64
	WriteMessage int64
}

var UserStatistics = make(map[int64]Statistics)
var MU sync.RWMutex
var lastUpdate time.Time

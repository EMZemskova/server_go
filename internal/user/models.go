package user

import "sync"

type User struct {
	ID       int64
	Username string
	Password string
}

type Statistics struct {
	ID           int64
	Username     string
	ChatsIn      int64
	WriteMessage int64
}

var userStatistics = make(map[int64]Statistics)
var mu sync.RWMutex

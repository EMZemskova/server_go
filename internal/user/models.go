package user

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

//var statsMap map[int64]Statistics
//statMap := make(map[int64]Statistics)

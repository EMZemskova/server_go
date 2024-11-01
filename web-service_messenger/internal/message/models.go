package message

type Message struct {
	ID     int64  `json:id`
	Chat   int64  `json:chat`
	Sender int64  `json:sender`
	Text   string `json:text`
}

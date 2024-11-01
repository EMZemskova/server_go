package chat

type ChatProvider interface {
	Create(chat Chat) (int, error)

	Get(id int64) (Chat, error)
}

package chat

type Provider interface {
	Create(chat Chat) (int, error)

	Get(id int64) (Chat, error)
}

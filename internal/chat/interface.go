package chat

type Provider interface {
	Create(chat Chat) (int, error)

	Get(id int64) (Chat, error)

	Edit(chat Chat, id int64) (Chat, error)

	Delete(chat Chat, id int64) (int, error)
}

package message

type MessageProvider interface {
	Create(message Message) (int, error)

	Get(id int64) (Message, error)

	Edit(message Message, id int64) (Message, error)

	Delete(id int64) error
}

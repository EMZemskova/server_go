package message

import (
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type message struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *message {
	return &message{db: db}
}

func (m *message) Create(message Message) (int, error) {
	query := `
	INSERT INTO messages (chat, sender,text)
	VALUES ($1, $2, $3)
	RETURNING id`
	var messageID int
	if err := m.db.QueryRow(query, message.Chat, message.Sender, message.Text).Scan(&messageID); err != nil {
		logrus.Error("failed create message", err)
		return 0, errors.Wrap(err, "failed to create message")
	}
	return messageID, nil
}

func (m *message) Get(id int64) (Message, error) {
	var findMessage Message
	query := `
        SELECT id, chat, sender,text
        FROM messages
        WHERE id = $1`
	if err := m.db.QueryRow(query, id).Scan(
		&findMessage.ID,
		&findMessage.Chat,
		&findMessage.Sender,
		&findMessage.Text,
	); err != nil {
		logrus.Error("getMessageById error: ", err)
		return Message{}, errors.Wrap(err, "getMessageById")
	}
	return findMessage, nil
}

func (m *message) Edit(message Message, id int64) (Message, error) {
	query := `
	UPDATE messages
	SET chat = $1, sender = $2, text = $3
	WHERE id = $4
	RETURNING id, chat, sender, text`
	if err := m.db.QueryRow(query,
		message.Chat,
		message.Sender,
		message.Text,
		id,
	).Scan(
		&message.ID,
		&message.Chat,
		&message.Sender,
		&message.Text,
	); err != nil {
		logrus.Error("failed to update message: ", err)
		return Message{}, errors.Wrap(err, "failed to update message")
	}
	return message, nil
}

func (m *message) Delete(id int64) error {
	query := `
        DELETE FROM messages
        WHERE id = $1`
	commandTag, err := m.db.Exec(query, id)
	if err != nil {
		logrus.Error("deleteMessage Error executing query:", err)
		return errors.Wrap(err, "failed to delete message")
	}
	if commandTag.RowsAffected() == 0 {
		logrus.Error("deleteMessage Error: no rows affected")
		return errors.New("no message found with the given id")
	}
	return nil
}

package message

import (
	"context"

	"github.com/jackc/pgx/v4"
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
	err := m.db.QueryRow(context.Background(),
		"INSERT INTO messages (chat, sender, text) VALUES ($1, $2, $3) RETURNING id", message.Chat, message.Sender, message.Text).Scan(&message.ID)
	if err != nil {
		logrus.Error("postMessage Error executing query:", err)
		return 0, errors.Wrap(err, "failed to create message")
	}
	id := int(message.ID)
	return id, err
}

func (m *message) Get(id int64) (Message, error) {
	var message Message
	err := m.db.QueryRow(context.Background(), "SELECT id, chat, sender, text FROM messages WHERE id=$1", id).
		Scan(&message.ID, &message.Chat, &message.Sender, &message.Text)
	if err != nil {
		logrus.Error("GetMessage Error executing query:", err)
		return Message{}, errors.Wrap(err, "failed to get message")
	}
	return message, nil
}

func (m *message) Edit(message Message, id int64) (Message, error) {

	err := m.db.QueryRow(context.Background(),
		"UPDATE messages SET chat=$1, sender=$2, text=$3 WHERE id=$4 RETURNING id, chat, sender, text",
		message.Chat, message.Sender, message.Text, id).Scan(&message.ID, &message.Chat, &message.Sender, &message.Text)
	if err != nil {
		logrus.Error("editMessage Error executing query:", err)
		return Message{}, errors.Wrap(err, "failed to edit message")
	}
	return message, nil
}

func (m *message) Delete(id int64) error {
	_, err := m.db.Exec(context.Background(), "DELETE FROM messages WHERE id=$1", id)
	if err != nil {
		logrus.Error("deleteMessage Error executing query:", err)
		return errors.Wrap(err, "failed to delete message")
	}
	return nil
}

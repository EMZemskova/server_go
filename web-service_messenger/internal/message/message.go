package message

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type message struct {
	db     *pgx.Conn
	id     int64
	chat   int64
	sender int64
	text   string
}

func (mess *message) Create(message Message) (int, error) {
	err := mess.db.QueryRow(context.Background(),
		"INSERT INTO messages (chat, sender, text) VALUES ($1, $2, $3) RETURNING id", message.Chat, message.Sender, message.Text).Scan(&message.ID)

	if err != nil {
		logrus.Error("postMessage Error executing query:", err)
		return 0, err
	}
	return int(message.ID), err
}

func (mess *message) Get(id int64) (Message, error) {
	var message Message

	err := mess.db.QueryRow(context.Background(), "SELECT id, chat, sender, text FROM messages WHERE id=$1", id).
		Scan(&message.ID, &message.Chat, &message.Sender, &message.Text)
	if err != nil {
		logrus.Error("GetMessage Error executing query:", err)
		return Message{}, err
	}

	return message, nil
}

func (mess *message) Edit(message Message, id int64) (Message, error) {

	err := mess.db.QueryRow(context.Background(),
		"UPDATE messages SET chat=$1, sender=$2, text=$3 WHERE id=$4 RETURNING id, chat, sender, text",
		message.Chat, message.Sender, message.Text, id).Scan(&message.ID, &message.Chat, &message.Sender, &message.Text)

	if err != nil {
		logrus.Error("editMessage Error executing query:", err)
		return Message{}, err
	}

	return message, nil
}

func (mess *message) Delete(id int64) error {
	_, err := mess.db.Exec(context.Background(), "DELETE FROM messages WHERE id=$1", id)
	if err != nil {
		logrus.Error("deleteMessage Error executing query:", err)
		return err
	}
	return nil
}

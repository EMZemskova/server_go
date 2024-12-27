package chat

import (
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type chat struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *chat {
	return &chat{db: db}
}

func (c *chat) Create(chat Chat) (int, error) {
	query := `
	INSERT INTO chats (creator, guest,status)
	VALUES ($1, $2, $3)
	RETURNING id`
	var chatID int
	if err := c.db.QueryRow(query, chat.Creator, chat.Guest, "Created").Scan(&chatID); err != nil {
		logrus.Error("failed create chat", err)
		return 0, errors.Wrap(err, "failed to create chat")
	}
	return chatID, nil
}

func (c *chat) Get(id int64) (Chat, error) {
	var findChat Chat
	query := `
        SELECT id, creator, guest, status
        FROM chats
        WHERE id = $1`
	if err := c.db.QueryRow(query, id).Scan(
		&findChat.ID,
		&findChat.Creator,
		&findChat.Guest,
		&findChat.Status,
	); err != nil {
		logrus.Error("getChatById error: ", err)
		return Chat{}, errors.Wrap(err, "getChatById")
	}
	return findChat, nil
}

func (c *chat) Edit(chat Chat, id int64) (Chat, error) {
	query := `
	UPDATE chats
	SET creator = $1, guest = $2, status = $3
	WHERE id = $4
	RETURNING id, creator, guest, status`
	if err := c.db.QueryRow(query,
		chat.Creator,
		chat.Guest,
		"Updated",
		id,
	).Scan(
		&chat.ID,
		&chat.Creator,
		&chat.Guest,
		&chat.Status,
	); err != nil {
		logrus.Error("failed to update chat: ", err)
		return Chat{}, errors.Wrap(err, "failed to update chat")
	}
	return chat, nil
}

func (c *chat) Delete(chat Chat, id int64) (int, error) {
	query := `
	UPDATE chats
	SET creator = $1, guest = $2, status = $3
	WHERE id = $4
	RETURNING id`
	if err := c.db.QueryRow(query,
		chat.Creator,
		chat.Guest,
		"Deleted",
		id,
	).Scan(
		&chat.ID,
	); err != nil {
		logrus.Error("failed to update chat: ", err)
		return 0, errors.Wrap(err, "failed to update chat")
	}
	return int(chat.ID), nil
}

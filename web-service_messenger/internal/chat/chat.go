package chat

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type chat struct {
	db *gorm.DB
}

func New(db *gorm.DB) *chat {
	return &chat{db: db}
}

func (c *chat) Create(chat Chat) (int, error) {

	result := c.db.Create(&chat)
	if result.Error != nil {
		logrus.Error("failed create chat", result.Error)
		return 0, errors.Wrap(result.Error, "failed to create chat")
	}
	return int(chat.ID), nil //?
}

func (c *chat) Get(id int64) (Chat, error) {

	var findchat Chat
	result := c.db.First(&findchat, id) //?
	if result.Error != nil {
		logrus.Error("getChatById", result.Error)
		return Chat{}, errors.Wrap(result.Error, "getChatById")
	}
	return findchat, nil
}

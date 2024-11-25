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
	chat.Status = "Created"
	result := c.db.Create(&chat)
	if result.Error != nil {
		logrus.Error("failed create chat", result.Error)
		return 0, errors.Wrap(result.Error, "failed to create chat")
	}
	return int(chat.ID), nil
}

func (c *chat) Get(id int64) (Chat, error) {
	var findchat Chat
	result := c.db.First(&findchat, id)
	if result.Error != nil {
		logrus.Error("getChatById", result.Error)
		return Chat{}, errors.Wrap(result.Error, "getChatById")
	}
	return findchat, nil
}

func (c *chat) Edit(chat Chat, id int64) (Chat, error) {
	chat.Status = "Updated"
	result := c.db.Model(&Chat{}).Where("id = ?", id).Updates(chat)
	if result.Error != nil {
		logrus.Error("failed update chat", result.Error)
		return Chat{}, errors.Wrap(result.Error, "failed to update chat")
	}
	return chat, nil
}

func (c *chat) Delete(chat Chat, id int64) (int, error) {
	chat.Status = "Deleted"
	result := c.db.Model(&Chat{}).Where("id = ?", id).Updates(chat)
	if result.Error != nil {
		logrus.Error("deleteChat Error:", result.Error)
		return 0, errors.Wrap(result.Error, "failed to delete chat")
	}
	return int(chat.ID), nil
}

package chat

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type chat struct {
	db      *gorm.DB
	id      int64
	creator int64
	guest   int64
}

func (ch *chat) Create(chat Chat) (int, error) {

	//need to replace Chat to chat?
	result := ch.db.Create(&chat)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to create chat")
	}
	return int(chat.ID), nil
}

func (ch *chat) Get(id int64) (Chat, error) {

	//chat or Chat or replace?
	var findchat Chat
	result := ch.db.First(&findchat, id)
	if result.Error != nil {
		logrus.Error(errors.Wrap(result.Error, "getChatById"))
		return findchat, result.Error //can't return nil?
	}
	return findchat, nil
}

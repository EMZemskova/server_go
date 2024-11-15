package message

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type message struct {
	db *gorm.DB
}

func New(db *gorm.DB) *message {
	return &message{db: db}
}

func (m *message) Create(message Message) (int, error) {
	result := m.db.Create(&message)
	if result.Error != nil {
		logrus.Error("failed create message", result.Error)
		return 0, errors.Wrap(result.Error, "failed to create message")
	}
	id := int(message.ID)
	return id, nil
}

func (m *message) Get(id int64) (Message, error) {
	var findmessage Message
	result := m.db.First(&findmessage, id)
	if result.Error != nil {
		logrus.Error("getMessageById", result.Error)
		return Message{}, errors.Wrap(result.Error, "getMessageById")
	}
	return findmessage, nil
}

func (m *message) Edit(message Message, id int64) (Message, error) {
	result := m.db.Model(&Message{}).Where("id = ?", id).Updates(message)
	if result.Error != nil {
		logrus.Error("failed update message", result.Error)
		return Message{}, errors.Wrap(result.Error, "failed to update message")
	}
	return message, nil
}

func (m *message) Delete(id int64) error {
	result := m.db.Delete(&Message{}, id)
	if result.Error != nil {
		logrus.Error("deleteMessage Error executing query:", result.Error)
		return errors.Wrap(result.Error, "failed to delete message")
	}
	return nil
}

package user

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func New(db *gorm.DB) *user {
	return &user{db: db}
}

func (u *user) Create(user User) (int, error) {

	result := u.db.Create(&user)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to create user")
	}
	return int(user.ID), nil //?
}

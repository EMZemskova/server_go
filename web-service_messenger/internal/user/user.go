package user

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type people struct {
	db        *gorm.DB
	id        int64
	username  string
	gpassword string
}

func (u *people) Create(user User) (int, error) {

	result := u.db.Create(&user)
	if result.Error != nil {
		return 0, errors.Wrap(result.Error, "failed to create user")
	}
	return int(user.ID), nil
}

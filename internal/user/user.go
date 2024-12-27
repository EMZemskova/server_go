package user

import (
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type user struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *user {
	return &user{db: db}
}

func (u *user) Create(user User) (int, error) {
	query := `
		INSERT INTO users (username, password)
		VALUES ($1, $2)
		RETURNING id`
	var userID int
	if err := u.db.QueryRow(query, user.Username, user.Password).Scan(&userID); err != nil {
		logrus.Error("failed create chat", err)
		return 0, errors.Wrap(err, "failed to create user")
	}
	return userID, nil
}

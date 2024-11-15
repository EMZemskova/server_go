package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	Gormdb *gorm.DB
	PgxDB  *pgx.Conn
}

func Init(connString string) (*storage, error) {
	DB := &storage{}
	var err error
	DB.Gormdb, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
		return nil, errors.Wrap(err, "failed to connect to the database")
	}
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
		return nil, errors.Wrap(err, "failed to connect to the database")
	}
	DB.PgxDB = conn
	logrus.Info("Database connection established successfully")
	return DB, nil
}

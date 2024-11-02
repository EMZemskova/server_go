package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	Gormdb *gorm.DB
	PgxDB  *pgx.Conn
}

func Init(connString string) (*storage, error) {
	var err error
	var DB storage
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
	log.Println("Database connection established successfully")
	return &DB, nil
}

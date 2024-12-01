package storage

import (
	"log"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	Gormdb *gorm.DB
}

func Init(connString string) (*storage, error) {
	db := &storage{}
	var err error
	db.Gormdb, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
		return nil, errors.Wrap(err, "failed to connect to the database")
	}
	return db, nil
}

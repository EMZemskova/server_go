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
	DB := &storage{}
	var err error
	DB.Gormdb, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
		return nil, errors.Wrap(err, "failed to connect to the database")
	}
	return DB, nil
}

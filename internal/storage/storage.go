package storage

import (
	"log"

	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
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

func (s *storage) RunMigrations(migrationsDir string) error {
	db, err := s.Gormdb.DB()
	if err != nil {
		return errors.Wrap(err, "failed to retrieve *sql.DB from Gorm")
	}
	if err := goose.SetDialect("postgres"); err != nil {
		return errors.Wrap(err, "failed to set Goose dialect")
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		return errors.Wrap(err, "failed to run migrations")
	}
	log.Println("Migrations applied successfully")
	return nil
}

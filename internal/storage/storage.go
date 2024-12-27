package storage

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

type storage struct {
	Conn *pgx.Conn
}

func Init(connString string) (*storage, error) {
	config, err := pgx.ParseConnectionString(connString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse connection string")
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	log.Println("Connected to database")
	return &storage{Conn: conn}, nil
}

func (s *storage) RunMigrations(migrationsDir string) error {
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return errors.Wrap(err, "failed to read migrations directory")
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}
		filePath := filepath.Join(migrationsDir, file.Name())
		log.Printf("Applying migration: %s", filePath)
		query, err := os.ReadFile(filePath)
		if err != nil {
			return errors.Wrapf(err, "failed to read migration file: %s", filePath)
		}
		if _, err := s.Conn.Exec(string(query)); err != nil {
			return errors.Wrapf(err, "failed to execute migration: %s", filePath)
		}
	}
	log.Println("All migrations applied successfully")
	return nil
}

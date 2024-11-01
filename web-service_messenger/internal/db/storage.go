package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	gormdb *gorm.DB
	pgxDB  *pgx.Conn
}

func Init(connString string) (*storage, error) {
	var err error
	var DB storage
	DB.gormdb, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
		return nil, err
	}

	pgxConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal("failed to parse config:", err)
		return nil, err
	}

	pgxPool, err1 := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err1 != nil {
		log.Fatal("failed to connect to the database:", err1)
		return nil, err
	}
	pgxConn, err2 := pgxPool.Acquire(context.Background())
	if err2 != nil {
		log.Fatal("failed to connect to the database:", err2)
		return nil, err
	}
	DB.pgxDB = pgxConn.Conn()
	log.Println("Database connection established successfully")
	return &DB, nil
}

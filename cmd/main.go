package main

import (
	"github.com/EMZemskova/server_go/internal"
	"github.com/EMZemskova/server_go/internal/chat"
	"github.com/EMZemskova/server_go/internal/handler"
	"github.com/EMZemskova/server_go/internal/message"
	metrics "github.com/EMZemskova/server_go/internal/prometheus"
	"github.com/EMZemskova/server_go/internal/stats"
	"github.com/EMZemskova/server_go/internal/storage"
	"github.com/EMZemskova/server_go/internal/user"
	"github.com/sirupsen/logrus"
)

func main() {
	connstring := "postgresql://postgres:postgres@postgres:5432/postgres"
	db, err := storage.Init(connstring)
	if err != nil {
		logrus.Fatal("Failed database connect", err)
	}
	migrationsDir := "./migrations"
	if err := db.RunMigrations(migrationsDir); err != nil {
		logrus.Fatalf("Error running migrations: %v", err)
	}

	userProvider := user.New(db.Conn)
	chatProvider := chat.New(db.Conn)
	messageProvider := message.New(db.Conn)
	statsProvider := stats.NewProvider(db.Conn)
	cacheStatsProvider := stats.NewCache(statsProvider)

	handle := handler.New(userProvider, chatProvider, messageProvider, cacheStatsProvider)
	router := internal.GetRouters(handle)
	metrics.InitMetrics("8081")
	logrus.Println("Metrics server started on port 8081")
	router.Run("0.0.0.0:8080")
	go func() {
		cacheStatsProvider.StartCacheUpdater()
	}()
}

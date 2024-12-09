package main

import (
	"github.com/EMZemskova/server_go/internal"
	"github.com/EMZemskova/server_go/internal/chat"
	"github.com/EMZemskova/server_go/internal/handler"
	"github.com/EMZemskova/server_go/internal/message"
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

	userProvider := user.New(db.Gormdb)
	chatProvider := chat.New(db.Gormdb)
	messageProvider := message.New(db.Gormdb)
	statsProvider := stats.NewProvider(db.Gormdb)
	cacheStatsProvider := stats.NewCache(statsProvider)

	handle := handler.New(userProvider, chatProvider, messageProvider, cacheStatsProvider)
	router := internal.GetRouters(handle)
	router.Run("0.0.0.0:8080")
	go func() {
		cacheStatsProvider.StartCacheUpdater()
	}()
}

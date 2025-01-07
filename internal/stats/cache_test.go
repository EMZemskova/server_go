package stats

import (
	"testing"
	"time"

	"github.com/EMZemskova/server_go/internal/storage"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	t.Run("NewCache initializes correctly", func(t *testing.T) {
		connstring := "postgresql://postgres:postgres@postgres:5432/postgres"
		db, err := storage.Init(connstring)
		if err != nil {
			logrus.Fatal("Failed database connect", err)
		}
		migrationsDir := "/app/migrations"
		if err := db.RunMigrations(migrationsDir); err != nil {
			logrus.Fatalf("Error running migrations: %v", err)
		}
		provider := NewProvider(db.Conn)
		cache := NewCache(provider)

		assert.NotNil(t, cache)
		assert.Equal(t, provider, cache.Provider)
		assert.NotNil(t, cache.userStatistics)
	})

	t.Run("StartCacheUpdater updates cache with real provider", func(t *testing.T) {
		connstring := "postgresql://postgres:postgres@postgres:5432/postgres"
		db, err := storage.Init(connstring)
		if err != nil {
			logrus.Fatal("Failed database connect", err)
		}
		migrationsDir := "/app/migrations"
		if err := db.RunMigrations(migrationsDir); err != nil {
			logrus.Fatalf("Error running migrations: %v", err)
		}
		provider := NewProvider(db.Conn)
		cache := NewCache(provider)
		go func() {
			cache.StartCacheUpdater()
		}()

		time.Sleep(35 * time.Second)
		cache.mu.RLock()
		defer cache.mu.RUnlock()
		assert.NotNil(t, cache.userStatistics)
		assert.True(t, len(cache.userStatistics) > 0, "cache.userStatistics should not be empty")
	})

	t.Run("StartCacheUpdater handles errors in real provider", func(t *testing.T) {
		connstring := "postgresql://postgres:postgres@postgres:5432/postgres"
		db, err := storage.Init(connstring)
		if err != nil {
			logrus.Fatal("Failed database connect", err)
		}
		migrationsDir := "/app/migrations"
		if err := db.RunMigrations(migrationsDir); err != nil {
			logrus.Fatalf("Error running migrations: %v", err)
		}
		provider := NewProvider(db.Conn)
		cache := NewCache(provider)

		go func() {
			cache.StartCacheUpdater()
		}()

		time.Sleep(35 * time.Second)

		cache.mu.RLock()
		defer cache.mu.RUnlock()
		assert.NotEmpty(t, cache.userStatistics, "cache.userStatistics should not be empty")
		assert.Greater(t, len(cache.userStatistics), 0, "cache.userStatistics should contain more than 0 elements")
	})
}

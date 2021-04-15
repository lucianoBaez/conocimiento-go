// Package repository Implements repository access for accounty
package repository

import (
	"context"
	"fmt"
	"github.com/d-Una-Interviews/svc_aut/internal/utils"
	"strings"
	"time"

	"github.com/Fs02/rel"
	"github.com/Fs02/rel/adapter/postgres"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // PostgreSQL migration
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // for migration
	"go.uber.org/zap"
)

const (
	// Default url for database scripts
	defaultSourceURL = "file://db/migrations"
)

var (
	logger, _   = zap.NewProduction(zap.Fields(zap.String("type", "repository")))
	dsnPostgres = ""
	shutdowns   []func() error
)

func init() {
	dsnPostgres = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		utils.GetEnv("POSTGRESQL_USERNAME", "user"),
		utils.GetEnv("POSTGRESQL_PASSWORD", "password"),
		utils.GetEnv("POSTGRESQL_HOST", "localhost"),
		utils.GetEnv("POSTGRESQL_PORT", "5432"),
		utils.GetEnv("POSTGRESQL_DATABASE", "drivers_db"))
}

// InitRepository Initialize Repositories.
// Connect to website database, run migrations and initialize the repository
func InitRepository() rel.Repository {

	adapter, err := postgres.Open(dsnPostgres)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
	// add to graceful shutdown list.
	shutdowns = append(shutdowns, adapter.Close)

	repository := rel.New(adapter)
	repository.Instrumentation(func(ctx context.Context, op string, message string) func(err error) {
		// no op for rel functions.
		if strings.HasPrefix(op, "rel-") {
			return func(error) {}
		}

		t := time.Now()

		return func(err error) {
			duration := time.Since(t)
			if err != nil {
				logger.Error(message, zap.Error(err), zap.Duration("duration", duration), zap.String("operation", op))
			} else {
				logger.Info(message, zap.Duration("duration", duration), zap.String("operation", op))
			}
		}
	})

	m, err := migrate.New(
		utils.GetEnv("SOURCE_URL", defaultSourceURL), dsnPostgres)
	if err != nil {
		logger.Info("uno")
		logger.Error(err.Error(), zap.Error(err))
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Info("dos")
		logger.Error(fmt.Sprintf("Unable to migrate up to the latest database schema - %v", err), zap.Error(err))
	}

	return repository
}

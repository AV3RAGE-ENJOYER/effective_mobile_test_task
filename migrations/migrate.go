package migrations

import (
	"log/slog"
	"os"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository/postgres"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose"
)

func RunMigrations(db *postgres.PostgresDB, ginMode string) {
	if err := goose.SetDialect("postgres"); err != nil {
		slog.Error("Failed to set postgres dialect", slog.Any("error", err))
		os.Exit(1)
	}

	driver := stdlib.OpenDBFromPool(db.Pool)

	slog.Info("Migrating database")

	migrationsDown := "migrations/debug"
	migrationsUp := "migrations/release"

	if ginMode == "debug" {
		migrationsDown = "migrations/release"
		migrationsUp = "migrations/debug"
	}

	if err := goose.Down(driver, migrationsDown); err != nil {
		slog.Error("Failed to drop migrations.", slog.Any("error", err), slog.String("mode", ginMode))
	}

	if err := goose.Up(driver, migrationsUp); err != nil {
		slog.Error("Failed to run migrations.", slog.Any("error", err), slog.String("mode", ginMode))
		os.Exit(1)
	}

	if err := driver.Close(); err != nil {
		slog.Error("Failed to close driver.", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("All migrations ran successfully!")
}

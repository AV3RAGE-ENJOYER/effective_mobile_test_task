package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/internal"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/migrations"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository/api"
	"github.com/AV3RAGE-ENJOYER/effective_mobile_test_task/repository/postgres"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           Music library test task
// @version         1.0
// @description     This is a test task for Juniour Go Developer in Effective Mobile.

// @contact.name   Andrei Dombrovskii
// @contact.email  andrushathegames@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8080
// @BasePath  /api/v1
func main() {
	// Setup logger

	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, handlerOpts))

	slog.SetDefault(logger)

	slog.Info("Logger is set up.")

	// Get enviromental variables

	err := godotenv.Load("config.env")

	if err != nil {
		slog.Error("Failed to load config.env file")
		os.Exit(1)
	}

	slog.Info("Successfully loaded config.env file.")

	GIN_MODE := os.Getenv("GIN_MODE")
	GIN_ADDR := os.Getenv("GIN_ADDR")
	POSTGRES_URL := os.Getenv("POSTGRES_URL")

	// Setup Database

	postgresDB, err := postgres.NewPostgresDB(context.Background(), POSTGRES_URL)

	if err != nil {
		slog.Error("Failed to establish connection to Postgres")
		os.Exit(1)
	}

	defer postgresDB.Pool.Close()

	slog.Info("Database is set up.")

	// Run Migrations

	migrations.RunMigrations(postgresDB, GIN_MODE)

	// Setup External API

	api := api.ExternalApiClient{
		BASE_URL: fmt.Sprintf("http://%s/api/v1", GIN_ADDR),
	}

	slog.Info("API is set up.")

	// Setup GIN

	slog.Info("GIN is set up.")

	gin.SetMode(GIN_MODE)

	serv := internal.SetupGinRouter(api, postgresDB)
	serv.Run(GIN_ADDR)
}

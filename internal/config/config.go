package config

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string
	Port  string
}

func Load(logger *slog.Logger) *Config {
	err := godotenv.Load()

	if err != nil {
		logger.Info("No .env file found, using system env")
	}

	return &Config{
		DBURL: getEnv("DB_URL", "postgres://user:pass@localhost:5432/db?sslmode=disable"),
		Port:  getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func ServerShutdown(srv *http.Server, logger *slog.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
	}
	logger.Info("Server exiting")
}

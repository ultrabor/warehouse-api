package config

import (
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
)

func RunMigrations(dbURL string) error {
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	slog.Info("Migrations applied successfully")
	return nil
}

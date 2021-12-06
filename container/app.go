package container

import (
	"github.com/secmohammed/restaurant-management/database"

	"github.com/secmohammed/restaurant-management/config"
)

type App struct {
	Config   *config.Config
	Database database.Database
}

func New(c *config.Config, d database.Database) *App {
	return &App{
		Config:   c,
		Database: d,
	}
}

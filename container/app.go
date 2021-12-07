package container

import (
	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/database"

	"github.com/secmohammed/restaurant-management/config"
)

type App struct {
	Config    *config.Config
	Database  database.Database
	Validator *validator.Validate
}

func New(c *config.Config, d database.Database, v *validator.Validate) *App {
	return &App{
		Config:    c,
		Database:  d,
		Validator: v,
	}
}

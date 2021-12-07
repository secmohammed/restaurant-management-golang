package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/secmohammed/restaurant-management/database"

	"github.com/secmohammed/restaurant-management/container"

	"github.com/secmohammed/restaurant-management/routes"

	"github.com/secmohammed/restaurant-management/config"
)

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	c := config.NewConfig()
	db := database.NewDatabaseConnection(c)
	validate := validator.New()

	// db := db.NewDB(c.DB)
	app := container.New(c, db, validate)
	r := routes.NewRouter(app)
	r.RegisterFoodRoutes()
	r.RegisterOrderRoutes()
	r.RegisterInvoiceRoutes()
	r.RegisterUserRoutes()
	r.RegisterMenuRoutes()
	r.RegisterOrderItemRoutes()
	r.RegisterTableRoutes()
	r.Serve()
}

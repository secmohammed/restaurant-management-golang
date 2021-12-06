package routes

import (
	"errors"
	"net/http"

	cors "github.com/rs/cors/wrapper/gin"
	"github.com/secmohammed/restaurant-management/container"

	"github.com/secmohammed/restaurant-management/utils"

	"github.com/gin-gonic/gin"
)

type Router interface {
	gin.IRouter
	Serve() error
	RegisterFoodRoutes()
	RegisterOrderRoutes()
	RegisterInvoiceRoutes()
	RegisterUserRoutes()
	RegisterMenuRoutes()
	RegisterOrderItemRoutes()
	RegisterTableRoutes()
}

type router struct {
	*gin.Engine
	app *container.App
}

func setupDefaults(r *gin.Engine) {
	// recover from error when server fails to start and retry.
	r.Use(gin.Recovery())
	r.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "OK"})
		return
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(utils.CreateApiError(http.StatusNotFound, errors.New("no route found")))
		return
	})
}

func NewRouter(app *container.App) Router {
	config := app.Config.Get()
	r := gin.New()
	if config.GetString("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.GetBool("app.log") {
		r.Use(gin.Logger())
	}
	r.Use(cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedHeaders: []string{"*"},
	}))
	setupDefaults(r)
	return &router{Engine: r, app: app}
}

func (r *router) Serve() error {
	port := r.app.Config.Get().GetString("app.port")
	return r.Run(":" + port)
}

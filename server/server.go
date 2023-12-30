package server

import (
	"example/httpserver/config"
	"example/httpserver/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Config *config.Configurations
}

func InitServer() *gin.Engine {
	router := gin.New()
	return router
}

func (app *App) InitRoutes() {
	app.Router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Health is ok")
	})

	{
		apiGroup := app.Router.Group("/api")
		eventHandler := controller.NewEventHandler()
		eventHandler.Config = app.Config
		apiGroup.POST("/event", eventHandler.PostEventInfoHandler)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zihangg/go-url-shortener/initializers"
	"github.com/zihangg/go-url-shortener/middleware"
	"github.com/zihangg/go-url-shortener/controllers"
)

func init() {
	initializers.LoadENV()
}

func main() {
	r := gin.Default()

	/* Defining middlewares */
	r.Use(middleware.ErrorHandler)

	/* Defining routes */
	r.GET("/health-check", controllers.HealthCheck)
	r.POST("/api/v1/shorten", controllers.Shorten)
	r.GET("/:url", controllers.Redirect)

	r.Run()
}
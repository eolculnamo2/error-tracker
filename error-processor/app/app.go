//Package appcontroller is the entry point for controllers
package app

import (
	"net/http"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


// StartRoutes initializes application routes
func StartRoutes() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"test": "test",})
	})
	router.Run()
}
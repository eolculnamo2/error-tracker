//Package appcontroller is the entry point for controllers
package appcontroller

import (
	"net/http"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/eolculnamo2/error-tracker/api/src/controllers/static_controller"
	"github.com/eolculnamo2/error-tracker/api/src/controllers/auth_controller"
	"github.com/eolculnamo2/error-tracker/api/src/controllers/org_controller"
)

const homePageFileLocation = "ui/build/index.html"

func homePage(c *gin.Context) {
		c.HTML(
				http.StatusOK,
				"index.html",
				gin.H{
						"title": "Home Page",
				},
		)
}

// StartRoutes initializes application routes
func StartRoutes() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Accept"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	//staticcontroller.StaticRoutes(router)
	authcontroller.AuthRoutes(router)
	orgcontroller.OrgRoutes(router)
	router.Run()
}
package staticcontroller

import (
	"net/http"
	"github.com/gin-gonic/gin"
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

// Initialize static routes.. i.e. HTML pages
func StaticRoutes(router *gin.Engine) {
	router.LoadHTMLFiles(homePageFileLocation)
	router.Static("/public/", "client/build/")
	router.GET("/", homePage)
}
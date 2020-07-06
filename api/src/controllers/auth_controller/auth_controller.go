
package authcontroller

import (
	"net/http"
	"github.com/eolculnamo2/error-tracker/api/src/services/new_user"
	"github.com/eolculnamo2/error-tracker/api/src/structs"
	"github.com/gin-gonic/gin"
)

	// c.JSON(http.StatusOK, gin.H{
	// 	"username": json.Email,
	// 	"password": json.Password,
	// })

const authRoute = "/auth"

func handleNewUser(c *gin.Context) {
	var json structs.NewUser

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newuser.CreateUser(json)
	c.Status(http.StatusOK)
}


func AuthRoutes(router *gin.Engine) {
	router.POST(authRoute + "/new-user", handleNewUser)
}

package authcontroller

import (
	"net/http"
	"github.com/eolculnamo2/error-tracker/api/src/services/new_user"
	"github.com/eolculnamo2/error-tracker/api/src/services/authentication"
	"github.com/eolculnamo2/error-tracker/api/src/structs"
	"github.com/eolculnamo2/error-tracker/api/src/helpers"
	"github.com/gin-gonic/gin"
)

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

func handleLogin(c *gin.Context) {
	var userLoginJson structs.UserLogin

	if err := c.ShouldBindJSON(&userLoginJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := authentication.AuthenticateUser(userLoginJson.Email, userLoginJson.Password)
	userDto := helpers.ConvertUserToDto(user)
	c.JSON(http.StatusOK, userDto)
}


func AuthRoutes(router *gin.Engine) {
	router.POST(authRoute + "/new-user", handleNewUser)
	router.POST(authRoute + "/login", handleLogin)
}
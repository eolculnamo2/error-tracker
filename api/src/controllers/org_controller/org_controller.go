package orgcontroller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/eolculnamo2/error-tracker/api/src/services/new_organization"
	"github.com/eolculnamo2/error-tracker/api/src/structs"
)


func newOrg(c *gin.Context) {
	var newOrgJson structs.NewOrg

	if err := c.ShouldBindJSON(&newOrgJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	neworganization.CreateNewOrg(newOrgJson)
	c.Status(http.StatusOK)
}

// Initialize static routes.. i.e. HTML pages
func OrgRoutes(router *gin.Engine) {
	router.POST("/new-org", newOrg)
}
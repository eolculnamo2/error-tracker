package resources

import (
	"github.com/eolculnamo2/error-tracker/error-processor/app/db/errordao"
	"github.com/eolculnamo2/error-tracker/error-processor/app/structs"
	"net/http"
	"github.com/gin-gonic/gin"
)

func handleErrorSubmit(c *gin.Context) {
	var webError structs.WebError

	if err := c.ShouldBindJSON(&webError); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	errordao.SaveNewWebError(webError)
}

func IncomingErrRoutes(router *gin.Engine) {
	router.POST("/submit-error", handleErrorSubmit)
}
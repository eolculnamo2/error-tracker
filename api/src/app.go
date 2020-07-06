package app

import (
	"github.com/jinzhu/gorm"
	"github.com/eolculnamo2/error-tracker/api/src/controllers"
	//"github.com/eolculnamo2/error-tracker/api/src/db"
)

const port = ":3000"

// Use microservices.. This is for user and organization data. Actual error data will process through another service
// If we don't want to spend money, then just deploy the other service on Heroku and not this one.

// StartApp begins the app
func StartApp(database *gorm.DB) {
	appcontroller.StartRoutes()
	//db.StartDb(database)
}
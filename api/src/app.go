package app

import (
	"github.com/jinzhu/gorm"
	"github.com/eolculnamo2/error-tracker/api/src/controllers"
	//"github.com/eolculnamo2/error-tracker/api/src/db"
)

// StartApp begins the app
func StartApp(database *gorm.DB) {
	appcontroller.StartRoutes()
	//db.StartDb(database)
}
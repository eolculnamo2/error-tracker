package userdb

import (
	"github.com/eolculnamo2/error-tracker/api/src/db"
	"github.com/eolculnamo2/error-tracker/api/src/models"
)

func FindUserByEmail(email string) models.User {
	var user models.User
	db.DbConnection.Where("email = ?", email).Find(&user)
	return user
}
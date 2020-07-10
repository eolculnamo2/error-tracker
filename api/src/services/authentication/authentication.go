package authentication

import (
	"github.com/eolculnamo2/error-tracker/api/src/models"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/eolculnamo2/error-tracker/api/src/db/user_db"
)


func AuthenticateUser(email string, password string) models.User {
	user := userdb.FindUserByEmail(email);
	err:= bcrypt.CompareHashAndPassword([]byte (user.Password), []byte (password))

	if err != nil {
		panic("Password does not match")
	}
	log.Println("User Data: ")
	log.Println(user.Email)
	log.Println(user.FirstName)
	return user
}
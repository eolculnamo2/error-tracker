// Package newuser is responsible for create a new user
package newuser

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/eolculnamo2/error-tracker/api/src/structs"
)

// bcrypt cost factor https://labs.clio.com/bcrypt-cost-factor-4ca0a9b03966
const cost = 15
// CreateUser creates a new user from the organization dashboard
func CreateUser(newUser structs.NewUser) {
	// hash password
	hashPw, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), cost)
	if err != nil {
		panic("Account creation failed: Bcrypt failed to hash password for new user: " + newUser.Email + " with error: " + err.Error())
	}
	newUser.Password = string(hashPw)
	fmt.Println(newUser.Password)
}
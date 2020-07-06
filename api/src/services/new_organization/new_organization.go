package neworganization

import (
	"github.com/eolculnamo2/error-tracker/api/src/structs"
	"github.com/eolculnamo2/error-tracker/api/src/db/organization_db"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewOrg(newOrg structs.NewOrg) {
	// newUser.Password = newOrg.Password
	hashPw, err := bcrypt.GenerateFromPassword([]byte(newOrg.Password), 12)
	if err != nil {
		panic("Account creation failed: Bcrypt failed to hash password for new user: " + newOrg.Email + " with error: " + err.Error())
	}
	newOrg.Password = string(hashPw)
	organizationdb.SaveOrg(newOrg)
}
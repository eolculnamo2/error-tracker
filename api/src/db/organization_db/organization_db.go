package organizationdb

import (
	"log"
	"github.com/eolculnamo2/error-tracker/api/src/models"
	"github.com/eolculnamo2/error-tracker/api/src/structs"
	"github.com/eolculnamo2/error-tracker/api/src/db"
	"github.com/eolculnamo2/error-tracker/api/src/constants/auth_types"
)

func SaveOrg(newOrg structs.NewOrg) {
	db.DbConnection.AutoMigrate(&models.Organization{}, &models.User{})
	var org = models.Organization{
		Name: newOrg.Name,
		Website: newOrg.Website,
		Enabled: 1,
		Users: []models.User{
			{
				Email: newOrg.Email,
				Password: newOrg.Password,
				FirstName: newOrg.FirstName,
				LastName: newOrg.LastName,
				Authority: authtypes.Admin,
				Enabled: 1,
			},
		},
	}
	db.DbConnection.Create(&org)
	log.Println("Created user: " + newOrg.Email)
}

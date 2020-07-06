package organizationdb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/eolculnamo2/error-tracker/api/src/models"
	"github.com/eolculnamo2/error-tracker/api/src/structs"
	"github.com/eolculnamo2/error-tracker/api/src/constants/auth_types"
)

// var db *gorm.DB

func SaveOrg(newOrg structs.NewOrg) {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)/error-tracker?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to the database" + err.Error())
	}
	db.AutoMigrate(&models.Organization{}, &models.User{})
	defer db.Close()
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
	db.Create(&org)
	//db.Save(&org)
	fmt.Println("Created user")
}

// func PassDb(database *gorm.DB) {
// 	orgStorage = OrgStorage{ db: database}
// }
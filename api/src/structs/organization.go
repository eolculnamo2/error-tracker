package structs

type NewOrg struct {
	Email        string `form:"email" json:"email" binding:"required"`  
	Password     string `form:"password" json:"password" binding:"required"`
	FirstName    string `form:"firstName" json:"firstName" binding:"required"`
	LastName     string `form:"lastName" json:"lastName" binding:"required"`
	Name     		 string `form:"name" json:"name" binding:"required"`
	Website      string `form:"website" json:"website" binding:"required"`
}
package structs

type NewUser struct {
	Email     		 string `form:"email" json:"email" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required"`
	OrganizationID int32 `form:"organizationId" json:"organizationId" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserModel without sensitive information like password
type UserDto struct {
		Email          string
		FirstName      string
		LastName       string
		OrganizationID uint
		Enabled        uint8
		Authority      string
}
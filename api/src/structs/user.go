package structs

type NewUser struct {
	Email     		 string `form:"email" json:"email" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required"`
	OrganizationID int32 `form:"organizationId" json:"organizationId" binding:"required"`
}
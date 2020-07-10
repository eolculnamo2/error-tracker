package helpers

import (
	"github.com/eolculnamo2/error-tracker/api/src/structs"
	"github.com/eolculnamo2/error-tracker/api/src/models"
)

func ConvertUserToDto(userModel models.User) structs.UserDto {
	var userDto structs.UserDto 
	userDto.Email = userModel.Email
	userDto.FirstName = userModel.FirstName
	userDto.LastName = userModel.LastName
	userDto.OrganizationID = userModel.OrganizationID
	userDto.Enabled = userModel.Enabled
	userDto.Authority = userModel.Authority
	return userDto
}
package models

type Organization struct {
	ID                uint
	Name              string
	Website           string 
	Enabled           uint8 
	Users 						[]User 
}

type User struct {
	Email        string
	Password     string 
	FirstName    string
	LastName     string
	OrganizationID uint
	Enabled      uint8
	Authority    string
}
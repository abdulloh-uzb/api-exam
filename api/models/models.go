package models

import (
	"api-exam/genproto/customer"
)

type Customer struct {
	FirstName    string
	LastName     string
	Email        string
	Username     string
	Password     string
	Bio          string
	PhoneNumber  string
	Addresses    []*customer.Address
	Code         string
	Refreshtoken string
}
type CustomerReq struct {
	FirstName   string
	LastName    string
	Email       string
	Username    string
	Password    string
	Bio         string
	PhoneNumber string
	Addresses   []*customer.Address
}
type UpdateCustomer struct {
	FirstName   string
	LastName    string
	Password    string
	PhoneNumber string
	Email       string
	Bio         string
	Addresses   []*customer.Address
}
type Address struct {
	Id       int64
	District string
	Street   string
}

type Verify struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	RefreshToken string
	AccessToken  string
}

type Login struct {
	Email        string
	FirstName    string
	LastName     string
	Password     string
	RefreshToken string
	AccessToken  string
}

type Admin struct {
	Email       string
	AccessToken string
}
type UpdatePost struct {
	Id          int64
	Name        string
	Description string
}
type UpdatePostResp struct {
	Id          int
	Name        string
	Description string
	CreatedAt   string
	UpdatedAt   string
	CustomerId  int
}

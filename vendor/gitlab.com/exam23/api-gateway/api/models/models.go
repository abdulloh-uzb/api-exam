package models

type Error struct {
	Code        int
	Error       error
	Description string
}

type Address struct {
	District   string `json:"district"`
	Street     string `json:"street"`
	HomeNumber string `json:"home_number"`
}

type CreateCustomer struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Bio         string    `json:"bio"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	Addresses   []Address `json:"addresses"`
}

type UpdateCustomer struct {
	Id          string          `json:"id"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	Email       string          `json:"email"`
	Bio         string          `json:"bio"`
	PhoneNumber string          `json:"phone_number"`
	Addresses   []UpdateAddress `json:"addresses"`
}

type UpdateAddress struct {
	Id         string `json:"id"`
	OwnerId    string `json:"owner_id"`
	District   string `json:"district"`
	Street     string `json:"street"`
	HomeNumber string `json:"home_number"`
}

type AddAddress struct {
	OwnerId    string `json:"owner_id"`
	District   string `json:"district"`
	Street     string `json:"street"`
	HomeNumber string `json:"home_number"`
}

type VerifyResponse struct {
	Id           string            `json:"id"`
	FirstName    string            `json:"first_name"`
	LastName     string            `json:"last_name"`
	Email        string            `json:"email"`
	Bio          string            `json:"bio"`
	PhoneNumber  string            `json:"phone_number"`
	JWT          string            `json:"jwt"`
	RefreshToken string            `json:"refresh"`
	Addresses    []AddressResponse `json:"addresses"`
}

type AddressResponse struct {
	Id         string `json:"id"`
	OwnerId    string `json:"owner_id"`
	District   string `json:"district"`
	Street     string `json:"street"`
	HomeNumber string `json:"home_number"`
}

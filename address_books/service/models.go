package service

import "time"

// AddressBook describe the struct required for storing diferent kind of locations and it's descriptions
type AddressBook struct {
	ID        int       `json:"id"`
	IsDefault bool      `json:"isDefault"`
	Name      string    `json:"name"`
	Note      string    `json:"note"`
	Address   Address   `json:"address"`
	Lat       string    `json:"lat"`
	Long      string    `json:"long"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Address describe the main details about the street, city, country and so on.
// The type is quite open for now any attribute can form part of the same, although might be validated after.
type Address map[string]interface{}

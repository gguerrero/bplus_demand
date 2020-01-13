package service

import (
	"github.com/google/uuid"
)

// Service interface defines the what the service layer should define
type Service interface {
	GetAddressBook(uuid.UUID) (AddressBook, error)
	CreateAddressBook(interface{}) (AddressBook, error)
	UpdateAddressBook(interface{}) (AddressBook, error)
	DeleteAddressBook(uuid.UUID) (AddressBook, error)
}

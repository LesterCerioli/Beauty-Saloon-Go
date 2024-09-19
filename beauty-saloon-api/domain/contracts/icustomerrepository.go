package contracts

import (
	"beauty-saloon-api/domain/models"
	"beauty-saloon-api/domain/valueobjects"
)


type ICustomerRepository interface {

	GetByCustomer(customerName string) (*models.Customer, error)


	GetByEmail(email string) (*models.Customer, error)


	GetByTelephone(telephone valueobjects.Telephone) (*models.Customer, error)


	GetList() ([]*models.Customer, error)


	Add(customer *models.Customer) error


	Update(customer *models.Customer) error


	Remove(customer *models.Customer) error
}



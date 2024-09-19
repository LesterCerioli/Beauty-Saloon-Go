package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	Email        string `json:"email"`
	Telephone    string `json:"telephone"`
}
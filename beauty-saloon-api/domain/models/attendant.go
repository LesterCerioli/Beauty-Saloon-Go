package models

import "gorm.io/gorm"

type Attendant struct {
	gorm.Model
	AttendantName string `json:"attendant_name"`
}
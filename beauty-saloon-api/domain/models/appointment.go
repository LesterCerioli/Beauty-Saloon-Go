package models

import (
	"time"

	"gorm.io/gorm"
)

// Appointment represents a salon appointment in the system.
type Appointment struct {
	gorm.Model
	Date            *time.Time   `json:"date"`
	AppointmentHour *time.Time   `json:"appointment_hour"`
	Attendants      []Attendant  `gorm:"many2many:appointment_attendants" json:"attendants"`
	Customer        Customer     `json:"customer"`
	Customers       []Customer   `gorm:"many2many:appointment_customers" json:"customers"`
}

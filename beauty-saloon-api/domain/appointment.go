package domain

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID              uuid.UUID
	ClientName      string
	AppointmentDate time.Time
	AppoitmentHour  time.Time
	AttendantName   string
	CustomerID      int
}

func NewAppointment(clientName string, appointmentDate time.Time, appointmentHour time.Time, attendantName string, customerID uuid.UUID) *Appointment {
	return &Appointment{
		ID:              uuid.New(), // Generate a new UUID for the ID
		ClientName:      clientName,
		AppointmentDate: appointmentDate,
		AppointmentHour: appointmentHour,
		AttendantName:   attendantName,
		CustomerID:      customerID,
	}
}

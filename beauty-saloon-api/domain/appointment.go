package domain

import "time"

type Appointment struct {
	ID             int
	ClientName     string
	AppointmentDate time.Time
	AttendantName  string
	CustomerID     int
}

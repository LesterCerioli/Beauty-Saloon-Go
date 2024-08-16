package repository

import (
	"beauty-saloon-api/domain"
	"database/sql"
)

type AppointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (repo *AppointmentRepository) Save(appointment *domain.Appointment) error {
	_, err := repo.db.Exec(`INSERT INTO Appointments (ClientName, AppointmentDate, AttendantName, CustomerID) VALUES (?, ?, ?, ?)`,
		appointment.ClientName, appointment.AppointmentDate, appointment.AttendantName, appointment.CustomerID)
	return err
}

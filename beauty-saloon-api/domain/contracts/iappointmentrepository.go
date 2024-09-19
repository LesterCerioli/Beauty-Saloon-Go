package contracts

import (
	"beauty-saloon-api/domain/models"
	"context"
	"time"
)

type IAppointmentRepository interface {
	GetByDate(ctx context.Context, date *time.Time) (*models.Appointment, error)

	GetByAppointmentHour(ctx context.Context, appointmentHour *time.Time) (*models.Appointment, error)

	GetList(ctx context.Context) ([]*models.Appointment, error)

	Add(ctx context.Context, appointment *models.Appointment) error

	Update(ctx context.Context, appointment *models.Appointment) error

	Remove(ctx context.Context, appointment *models.Appointment) error
}

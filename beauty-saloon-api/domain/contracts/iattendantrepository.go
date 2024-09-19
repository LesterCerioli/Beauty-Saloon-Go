package contracts

import (
	"beauty-saloon-api/domain/models"
	"context"
)

// AttendantRepository defines the interface for managing attendants in the system.
type AttendantRepository interface {
	GetByAttendantName(ctx context.Context, attendantName string) (*models.Attendant, error)
	GetList(ctx context.Context) ([]models.Attendant, error)
	Add(ctx context.Context, attendant *models.Attendant) error
	Update(ctx context.Context, attendant *models.Attendant) error
	Remove(ctx context.Context, attendant *models.Attendant) error
}
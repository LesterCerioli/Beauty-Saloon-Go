package contracts

import (
	"context"
	"beauty-saloon-api/domain/models"
)

// IStateRepository defines the methods for interacting with the State data.
type IStateRepository interface {
	GetByStateName(ctx context.Context, stateName string) (*models.State, error)
	GetByUF(ctx context.Context, uf string) (*models.State, error)
	GetList(ctx context.Context) ([]*models.State, error)
	Add(ctx context.Context, state *models.State) error
	Update(ctx context.Context, state *models.State) error
	Remove(ctx context.Context, state *models.State) error
}

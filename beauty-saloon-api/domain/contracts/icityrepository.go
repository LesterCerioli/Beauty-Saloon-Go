package contracts

import (
	"context"
	"beauty-saloon-api/domain/models"
)


type ICityRepository interface {
	GetByCityName(ctx context.Context, cityName string) (*models.City, error)
	GetList(ctx context.Context) ([]*models.City, error)
	Add(ctx context.Context, city *models.City) error
	Update(ctx context.Context, city *models.City) error
	Remove(ctx context.Context, city *models.City) error
}

package contracts

import (
	"context"
	"beauty-saloon-api/domain/models"
	"beauty-saloon-api/domain/valueobjects"
)

type ISaloonRepository interface {
	
	GetByFantasyName(ctx context.Context, fantasyName string) (*models.Saloon, error)

	
	GetByAddress(ctx context.Context, address valueobjects.Address) (*models.Saloon, error)

	
	GetByDistrict(ctx context.Context, district valueobjects.District) (*models.Saloon, error)

	
	GetList(ctx context.Context) ([]*models.Saloon, error)

	
	Add(ctx context.Context, saloon *models.Saloon) error

	
	Update(ctx context.Context, saloon *models.Saloon) error

	
	Remove(ctx context.Context, saloon *models.Saloon) error
}



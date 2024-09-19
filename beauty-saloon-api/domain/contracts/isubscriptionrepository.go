package contracts

import (
	"beauty-saloon-api/domain/models"
	"beauty-saloon-api/domain/valueobjects"
)

type SubscriptionRepository interface {

	GetByCnpj(cnpj valueobjects.Cnpj) (*models.Subscription, error)
	

	GetBySubscriptionNumber(subscriptionNumber string) (*models.Subscription, error)
	

	Add(subscription *models.Subscription) error
	

	Update(subscription *models.Subscription) error
	

	Remove(subscription *models.Subscription) error
}

package main

import (
	"fmt"
	"net/http"
	"beauty-saloon-api/domain/models"
	"beauty-saloon-api/domain/valueobjects"
)


type SubscriptionRepositoryImpl struct {
	subscriptions []models.Subscription
}

func (repo *SubscriptionRepositoryImpl) GetByCnpj(cnpj valueobjects.Cnpj) (*models.Subscription, error) {
	for _, subscription := range repo.subscriptions {
		if subscription.Cnpj == cnpj {
			return &subscription, nil
		}
	}
	return nil, fmt.Errorf("subscription not found")
}

func (repo *SubscriptionRepositoryImpl) GetBySubscriptionNumber(subscriptionNumber string) (*models.Subscription, error) {
	for _, subscription := range repo.subscriptions {
		if subscription.SubscriptionNumber == subscriptionNumber {
			return &subscription, nil
		}
	}
	return nil, fmt.Errorf("subscription not found")
}

func (repo *SubscriptionRepositoryImpl) Add(subscription *models.Subscription) error {
	repo.subscriptions = append(repo.subscriptions, *subscription)
	return nil
}

func (repo *SubscriptionRepositoryImpl) Update(subscription *models.Subscription) error {
	for i, s := range repo.subscriptions {
		if s.SubscriptionNumber == subscription.SubscriptionNumber {
			repo.subscriptions[i] = *subscription
			return nil
		}
	}
	return fmt.Errorf("subscription not found")
}

func (repo *SubscriptionRepositoryImpl) Remove(subscription *models.Subscription) error {
	for i, s := range repo.subscriptions {
		if s.SubscriptionNumber == subscription.SubscriptionNumber {
			repo.subscriptions = append(repo.subscriptions[:i], repo.subscriptions[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("subscription not found")
}

func subscriptionHandler(w http.ResponseWriter, r *http.Request) {
	repo := &SubscriptionRepositoryImpl{}
	
	
	cnpj, err := valueobjects.NewCnpj("12.345.678/0001-23")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subscription := models.NewSubscription(cnpj, "123456789")

	
	err = repo.Add(subscription)
	if err != nil {
		http.Error(w, "Failed to add subscription", http.StatusInternalServerError)
		return
	}

	
	result, err := repo.GetByCnpj(cnpj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	
	fmt.Fprintf(w, "Subscription found: %v", result)
}

func main() {
	
	http.HandleFunc("/subscription", subscriptionHandler)

	
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

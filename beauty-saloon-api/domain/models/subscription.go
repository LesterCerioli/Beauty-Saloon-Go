package models

import (
	"math/rand"
	"strings"
	"time"
	"beauty-saloon-api/domain/valueobjects" // Importando o value object Cnpj
)

type Subscription struct {
	Cnpj              valueobjects.Cnpj `json:"cnpj"`               // Cnpj como Value Object
	SubscriptionNumber string            `json:"subscription_number"`
	Saloons            []Saloon          `json:"saloons"`
}

// NewSubscription cria uma nova assinatura com um CNPJ e número de assinatura
func NewSubscription(cnpj valueobjects.Cnpj, subscriptionNumber string) *Subscription {
	if subscriptionNumber == "" {
		subscriptionNumber = GenerateSubscriptionNumber()
	}

	return &Subscription{
		Cnpj:              cnpj,
		SubscriptionNumber: subscriptionNumber,
		Saloons:            make([]Saloon, 0),
	}
}


func GenerateSubscriptionNumber() string {
	vowels := "AEIOU"
	numbers := "0123456789"
	rand.Seed(time.Now().UnixNano())

	var randomVowels strings.Builder
	for i := 0; i < 3; i++ {
		randomVowels.WriteByte(vowels[rand.Intn(len(vowels))])
	}

	var randomNumbers strings.Builder
	for i := 0; i < 7; i++ {
		randomNumbers.WriteByte(numbers[rand.Intn(len(numbers))])
	}

	return randomVowels.String() + randomNumbers.String()
}

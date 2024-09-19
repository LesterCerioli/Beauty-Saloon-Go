package models

import "gorm.io/gorm"

type Saloon struct {
	gorm.Model
	FantasyName  string  `json:"fantasy_name"`
	SocialReason string  `json:"social_reason"`
	Cnpj         string  `json:"cnpj"`
	OwnerName    string  `json:"owner_name"`
	Telephone    string  `json:"telephone"`
	Address      string  `json:"address"`
	District     string  `json:"district"`
}
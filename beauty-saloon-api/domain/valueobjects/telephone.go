package valueobjects

import (
	"errors"
	"fmt"
	"regexp"
	"strings" // Correct import for string manipulation
)

type Telephone struct {
	TelephoneNumber string
	TelephoneRegion string
}

func NewTelephone(telephoneNumber, telephoneRegion string) (*Telephone, error) {
	if !ValidateTelephone(telephoneNumber) {
		return nil, errors.New("invalid telephone number")
	}
	return &Telephone{
		TelephoneNumber: telephoneNumber,
		TelephoneRegion: telephoneRegion,
	}, nil
}

// ValidateTelephone validates the telephone number
func ValidateTelephone(telephone string) bool {
	// Trims whitespace
	shortenNum := strings.TrimSpace(telephone)

	// Removes non-alphanumeric characters
	shortenNum = regexp.MustCompile(`[^0-9a-zA-Z]+`).ReplaceAllString(shortenNum, "")

	// Checks if the cleaned number has 13 characters
	if len(shortenNum) == 13 {
		fmt.Println("O número de telefone é válido")
		return true
	}
	fmt.Println("O número de telefone é inválido")
	return false
}

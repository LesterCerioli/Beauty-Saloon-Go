package valueobjects

import (
	"fmt"
	"regexp"
	"strings"
)


type Cnpj string


func NewCnpj(cnpjStr string) (Cnpj, error) {
	
	cnpjStr = sanitizeCnpj(cnpjStr)

	
	if !isValidCnpj(cnpjStr) {
		return "", fmt.Errorf("CNPJ inválido: %s", cnpjStr)
	}

	return Cnpj(cnpjStr), nil
}


func sanitizeCnpj(cnpjStr string) string {
	return strings.ReplaceAll(strings.ReplaceAll(cnpjStr, ".", ""), "-", "")
}


func isValidCnpj(cnpjStr string) bool {
	
	if len(cnpjStr) != 14 {
		return false
	}

	
	match, _ := regexp.MatchString(`^\d{14}$`, cnpjStr)
	return match
}


func (c Cnpj) FormatCnpj() string {
	formatted := fmt.Sprintf("%s.%s.%s/%s-%s",
		c[:2], c[2:5], c[5:8], c[8:12], c[12:14])
	return formatted
}


func (c Cnpj) String() string {
	return string(c)
}

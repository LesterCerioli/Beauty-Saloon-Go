package valueobjects

type Address struct {
	AvenueOrStreet string
	Number         string
}


func NewAddress(avenueOrStreet, number string) *Address {
	return &Address{
		AvenueOrStreet: avenueOrStreet,
		Number:         number,
	}
}

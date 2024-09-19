package valueobjects

type District struct {
	DistrictName string
}

func NewDistrict(districtName string) *District {
	return &District{
		DistrictName: districtName,
	}
}

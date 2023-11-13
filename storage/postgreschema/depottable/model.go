package depottable

type Depot struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	CityId       int64  `json:"cityId"`
	CityName     string `json:"cityName"`
	DistrictId   int64  `json:"districtId"`
	DistrictName string `json:"districtName"`
	WardId       int64  `json:"wardId"`
	WardName     string `json:"wardName"`
	Address      string `json:"address"`
}

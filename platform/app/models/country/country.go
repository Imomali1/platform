package country

type Country struct {
	CountryId   int    `json:"country_id"`
	CountryName string `json:"country_name"`
	RegionId    int    `json:"region_id"`
}

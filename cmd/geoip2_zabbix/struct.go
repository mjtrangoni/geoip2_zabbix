package main

type GeoIPCityJSON struct {
	City struct {
		Name string `json:"name"`
	} `json:"city"`
	Continent struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"continent"`
	Country struct {
		IsInEuropeanUnion bool   `json:"is_in_european_union"`
		IsoCode           string `json:"iso_code"`
		Name              string `json:"name"`
	} `json:"country"`
	Location struct {
		AccuracyRadius uint    `json:"accuracy_radius"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		MetroCode      uint    `json:"metro_code"`
		TimeZone       string  `json:"time_zone"`
	} `json:"location"`
	PostalCode string `json:"postal_code"`
	Traits     struct {
		IsAnonymousProxy    bool `json:"is_anonymous_proxy"`
		IsSatelliteProvider bool `json:"is_satellite_provider"`
	} `json:"traits"`
}

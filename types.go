package main

// I know not everyone like the "types.go" file, but I do.
// So, here it is.

type Person struct {
	Name       string `json:"name"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Relation   string `json:"relation"`
	Id         string `json:"id"`
}

type Location struct {
	IsActive     bool   `json:"is_active"`
	StreetLine1  string `json:"street_line_1"`
	StreetLine2  string `json:"street_line_2"`
	City         string `json:"city"`
	StateCode    string `json:"state_code"`
	PostalCode   string `json:"postal_code"`
	ZIP4         string `json:"zip4"`
	CountryCode  string `json:"country_code"`
	LocationType string `json:"location_type"`
	LatLong      struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Accuracy  string  `json:"accuracy"`
	} `json:"lat_long"`
	LinkToPersonStartDate string `json:"link_to_person_start_date"`
	LinkToPersonEndDate   string `json:"link_to_person_end_date"`
	DeliveryPoint         string `json:"delivery_point"`
	Id                    string `json:"id"`
}

type TrestleResponse struct {
}

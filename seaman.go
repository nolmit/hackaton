package main

// Vessel stores a value
type Seaman struct {
	ID string `json:"id"`
	Rank string `json:"rank"`
	Name string `json:"name"`
	Nationality string `json:"nationality"`
	DateOfBirth string `json:"dateOfBirth"`
	PlaceofBirth string `json:"placeOfBirth"`
	PassportNumber string `json:"passportNumber"`
	PassportExpiry string `json:"passportExpiry"`
	SeamanBookNumber string `json:"seamanBookNumber"`
	SeamanBookExpiry string `json:"seamanBookExpiry"`
}
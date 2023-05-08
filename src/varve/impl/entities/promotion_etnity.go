package entities

import "encoding/json"

type Promotion struct {
	ID             string `json:"id"`
	Price          string `json:"price"`
	ExpirationDate string `json:"expiration_date"`
}

func (p Promotion) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

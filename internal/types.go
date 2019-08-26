package internal

import "time"

// struct to unmarshal API response into
type APIResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ValueExcVat float64   `json:"value_exc_vat"`
		ValueIncVat float64   `json:"value_inc_vat"`
		ValidFrom   time.Time `json:"valid_from"`
		ValidTo     time.Time `json:"valid_to"`
	} `json:"results"`
}

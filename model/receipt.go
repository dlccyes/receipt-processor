package model

import "time"

type Receipt struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate time.Time `json:"purchaseDate"`
	PurchaseTime time.Time `json:"purchaseTime"`
	Items        []Item    `json:"items"`
	Total        float64   `json:"total"`
}

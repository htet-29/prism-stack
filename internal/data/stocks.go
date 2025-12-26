package data

import "time"

type Stock struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ItemName  string    `json:"item_name"`
	SKU       string    `json:"sku"`
	Category  []string  `json:"category"`
	UnitPrice int64     `js:"unit_price"`
	Quantity  int64     `json:"quantity"`
	Version   int64     `json:"version"`
}

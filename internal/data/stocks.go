package data

import "time"

type Stock struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	ItemName  string    `json:"item_name"`
	SKU       string    `json:"sku"`
	Category  []string  `json:"category,omitzero"`
	UnitPrice int64     `json:"unit_price,omitzero"`
	Quantity  int64     `json:"quantity,omitzero"`
	Version   int64     `json:"version"`
}

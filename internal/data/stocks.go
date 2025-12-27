package data

import (
	"time"

	"github.com/htet-29/prism-stack/internal/validator"
)

type Stock struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	ItemName  string    `json:"item_name"`
	SKU       string    `json:"sku"`
	Category  []string  `json:"category,omitzero"`
	UnitPrice Unit      `json:"unit_price,omitzero"`
	Quantity  int64     `json:"quantity,omitzero"`
	Version   int64     `json:"version"`
}

func ValidateStock(v *validator.Validator, stock *Stock) {
	v.Check(stock.ItemName != "", "item_name", "must be provided")
	v.Check(len(stock.ItemName) <= 500, "item_name", "must not be more than 500 bytes long")

	v.Check(stock.SKU != "", "sku", "must be provided")
	// TODO: check if sku is unique and max bytes

	v.Check(stock.Category != nil, "category", "must be provided")
	v.Check(len(stock.Category) >= 1, "category", "must contain at least 1 category")
	// TODO: check if max number of categories

	v.Check(stock.UnitPrice > 0, "unit_price", "must be greater than zero")
	v.Check(stock.Quantity > 0, "quantity", "must be greater than zero")

	v.Check(validator.Unique(stock.Category), "category", "must not contain duplicate values")
}

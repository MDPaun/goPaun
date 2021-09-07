package models

import (
	"database/sql"
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Inventory struct {
	ID       int     `json:"id"`
	Image    string  `json:"image"`
	Name     string  `json:"name"`
	SKU      string  `json:"sku"`
	EAN      string  `json:"ean"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type FilterProducts struct {
	PageNo     int    `json:"page"`
	PageLimit  int    `json:"page_limit"`
	SortName   string `json:"sort_name"`
	SortSKU    string `json:"sort_sku"`
	SortEAN    string `json:"sort_ean"`
	SortOnHand string `json:"sort_on_hand"`
}
type InventoryModel struct {
	DBDC *sql.DB
	DBMC *sql.DB
	DB   *sql.DB
}

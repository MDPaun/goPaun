package models

import (
	"database/sql"
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Inventory struct {
	ID       int    `json:"id"`
	Image    string `json:"image"`
	Name     string `json:"name"`
	SKU      string `json:"sku"`
	EAN      string `json:"ean"`
	Quantity int    `json:"quantity"`
}

type InventoryModel struct {
	DBDC *sql.DB
	DB   *sql.DB
}

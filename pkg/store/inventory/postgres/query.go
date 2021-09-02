package inventory

import (
	"database/sql"
	"errors"
	"fmt"

	models "github.com/MDPaun/goPaun/pkg/store/inventory"
)

type InventoryModel struct {
	DB *sql.DB
}

// This will return a specific staff member based on its id.
func (m *InventoryModel) GetBySKU(sku string) (*models.Inventory, error) {

	stmt := "SELECT id, image, name, sku, ean, quantity FROM products WHERE sku = $1;"

	row := m.DB.QueryRow(stmt, sku)
	s := &models.Inventory{}
	err := row.Scan(&s.ID, &s.Image, &s.Name, &s.SKU, &s.EAN, &s.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *InventoryModel) GetByID(id int) (*models.Inventory, error) {

	stmt := "SELECT id, image, name, sku, ean, quantity FROM products WHERE sku = $1;"

	row := m.DB.QueryRow(stmt, id)
	s := &models.Inventory{}
	err := row.Scan(&s.ID, &s.Image, &s.Name, &s.SKU, &s.EAN, &s.Quantity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 10 most recently created members.
func (m *InventoryModel) Latest(page int) ([]*models.Inventory, error) {
	stmt := "SELECT id, image, name, sku, ean, quantity FROM products ORDER BY id ASC LIMIT $1 OFFSET $2;"
	DefaultLimit := 10

	rows, err := m.DB.Query(stmt, DefaultLimit, (page-1)*DefaultLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	inventory := []*models.Inventory{}

	for rows.Next() {
		s := &models.Inventory{}
		err = rows.Scan(&s.ID, &s.Image, &s.Name, &s.SKU, &s.EAN, &s.Quantity)
		if err != nil {
			return nil, err
		}
		inventory = append(inventory, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return inventory, nil
}

func (m *InventoryModel) AddProduct(image, name, sku, ean string, quantity int) error {
	// fmt.Println(id, stock)
	stmt := "INSERT INTO products (image, name, sku, ean, quantity) VALUES($1, $2, $3, $4, $5)"

	_, err := m.DB.Exec(stmt, image, name, sku, ean, quantity)
	if err != nil {
		return err
	}

	return nil
}

func (m *InventoryModel) UpdateStock(sku, quantity string) error {

	stmt := "UPDATE products SET quantity = $1  WHERE sku = $2;"

	_, err := m.DB.Exec(stmt, quantity, sku)
	if err != nil {
		return err
	}

	return nil
}

func (m *InventoryModel) CountProduct() {

	stmt := "select count (*) from products"
	// var s int64
	row, _ := m.DB.Query(stmt)
	fmt.Println(row)

}

// func (m *InventoryModel) UpdateStock() ([]*models.Inventory, error) {
// 	stmt := "SELECT id, image, name, sku, ean, quantity FROM products LIMIT 10;"
// 	rows, err := m.DB.Query(stmt)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	inventory := []*models.Inventory{}

// 	for rows.Next() {
// 		s := &models.Inventory{}
// 		err = rows.Scan(&s.ID, &s.Image, &s.Name, &s.SKU, &s.EAN, &s.Quantity)
// 		if err != nil {
// 			return nil, err
// 		}
// 		inventory = append(inventory, s)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return inventory, nil
// }

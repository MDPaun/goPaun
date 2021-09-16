package inventory

import (
	"database/sql"
	"errors"
	"log"

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
func (m *InventoryModel) Latest(page, defaultLimit int, sortName, sortSKU, sortEAN, sortOnHand string) ([]*models.Inventory, error) {

	stmt := `SELECT id, image, name, sku, ean, quantity, price FROM products 
				WHERE
					name ILIKE '%'||$1||'%' AND
					sku ILIKE '%'||$2||'%' AND
					ean ILIKE '%'||$3||'%' AND
					quantity::text ILIKE '%'||$4||'%' AND
					price > 0
				ORDER BY id ASC  LIMIT $5 OFFSET $6;`

	rows, err := m.DB.Query(stmt, sortName, sortSKU, sortEAN, sortOnHand, defaultLimit, (page-1)*defaultLimit)
	// rows, err := m.DB.Query(stmt)
	// rows, err := m.DB.Query.SELECT("asdasd")

	// rows, err := m.DB.Query(stmt, sortName, (page-1)*defaultLimit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	inventory := []*models.Inventory{}

	for rows.Next() {
		s := &models.Inventory{}
		err = rows.Scan(&s.ID, &s.Image, &s.Name, &s.SKU, &s.EAN, &s.Quantity, &s.Price)
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

func (m *InventoryModel) AddProduct(image, name, sku, ean string, quantity int, price float64) error {
	// fmt.Println(id, stock)
	stmt := "INSERT INTO products (image, name, sku, ean, quantity, price) VALUES($1, $2, $3, $4, $5, $6)"

	_, err := m.DB.Exec(stmt, image, name, sku, ean, quantity, price)
	if err != nil {
		return err
	}

	return nil
}

func (m *InventoryModel) UpdateStock(sku, quantity string, price float64) error {

	stmt := "UPDATE products SET quantity = $1, price = $2  WHERE sku = $3;"

	_, err := m.DB.Exec(stmt, quantity, price, sku)
	if err != nil {
		return err
	}

	return nil
}

func (m *InventoryModel) CountProduct() (total int) {
	var count int
	// stmt := "SELECT COUNT (*) as count from products"
	stmt, err := m.DB.Prepare("SELECT COUNT(*) as count FROM products")
	if err != nil {
		log.Fatal(err)
	}

	// m.DB.QueryRow(stmt).Scan(&count)
	err = stmt.QueryRow().Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(count)
	stmt.Close()
	return count
}

func (m *InventoryModel) GetAllSK() ([]*models.Inventory, error) {

	stmt := "SELECT sku, ean, price FROM products WHERE sku ILIKE 'SK%' AND COALESCE(ean, '') <> '';"

	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	inventory := []*models.Inventory{}

	for rows.Next() {
		s := &models.Inventory{}
		err = rows.Scan(&s.SKU, &s.EAN, &s.Price)
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

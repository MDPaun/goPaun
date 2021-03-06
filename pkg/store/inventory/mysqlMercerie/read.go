package mysqlMercerie

import (
	"database/sql"
	"errors"

	models "github.com/MDPaun/goPaun/pkg/store/inventory"
)

type InventoryModel struct {
	DBMC *sql.DB
}

// This will return a specific staff member based on its id.
func (m *InventoryModel) FindByID(id int) (*models.Inventory, error) {

	stmt := "SELECT id, email, fullname, image, status, date_added FROM staff WHERE id = $1;"

	row := m.DBMC.QueryRow(stmt, id)
	s := &models.Inventory{}
	err := row.Scan(&s.ID, &s.SKU, &s.EAN, &s.Quantity)
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
func (m *InventoryModel) GetProducts() ([]*models.Inventory, error) {
	stmt := `SELECT product.product_id, product.sku, product.ean, product.image, product.quantity, product.price,product_description.name
				FROM product
				INNER JOIN product_description ON product.product_id = product_description.product_id
				LIMIT 10`
	rows, err := m.DBMC.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	inventory := []*models.Inventory{}

	for rows.Next() {
		s := &models.Inventory{}
		err = rows.Scan(&s.ID, &s.SKU, &s.EAN, &s.Image, &s.Quantity, &s.Price, &s.Name)
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

func (m *InventoryModel) UpdateStockMercerie(sku, quantity string, price float64) error {
	// fmt.Println(id, stock)
	stmt := "UPDATE product SET product.quantity = ?, price = ?  WHERE sku = ?;"

	_, err := m.DBMC.Exec(stmt, quantity, price, sku)
	if err != nil {
		return err
	}

	return nil
}

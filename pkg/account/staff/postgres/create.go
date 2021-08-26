package postgres

import (
	// "database/sql"

	"time"
)

func (m *StaffModel) Create(email, fullname, image string, status bool, date_added time.Time) error {

	stmt := "INSERT INTO staff (email, fullname, image, status, date_added) VALUES($1, $2, $3, $4, $5)"

	_, err := m.DB.Exec(stmt, email, fullname, image, status, date_added)
	if err != nil {
		return err
	}

	return nil
}

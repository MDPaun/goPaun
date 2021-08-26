package postgres

import (
	// "database/sql"

	"database/sql"
	"errors"

	models "github.com/MDPaun/goPaun/pkg/account/staff"
)

type StaffModel struct {
	DB *sql.DB
}

// This will return a specific staff member based on its id.
func (m *StaffModel) FindByID(id int) (*models.Staff, error) {
	stmt := "SELECT id, email, fullname, image, status, date_added FROM staff WHERE id = $1;"

	row := m.DB.QueryRow(stmt, id)
	s := &models.Staff{}
	err := row.Scan(&s.ID, &s.Email, &s.FullName, &s.Image, &s.Status, &s.DateAdded)
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
func (m *StaffModel) Latest() ([]*models.Staff, error) {
	stmt := "SELECT id, email, fullname, image, status, date_added  FROM staff LIMIT 10;"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	staff := []*models.Staff{}

	for rows.Next() {
		s := &models.Staff{}
		err = rows.Scan(&s.ID, &s.Email, &s.FullName, &s.Image, &s.Status, &s.DateAdded)
		if err != nil {
			return nil, err
		}
		staff = append(staff, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return staff, nil
}

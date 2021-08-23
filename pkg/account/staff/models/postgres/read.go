package postgres

import (
	// "database/sql"

	"database/sql"

	"github.com/MDPaun/goPaun/pkg/account/staff/models"
)

type StaffModel struct {
	DB *sql.DB
}

func (m *StaffModel) Create(title, content, expires string) (int, error) {
	return 0, nil
}

// This will return a specific staff member based on its id.
func (m *StaffModel) Read(id int) (*models.Staff, error) {
	return nil, nil
}

// This will return the 10 most recently created members.
func (m *StaffModel) Latest() ([]*models.Staff, error) {
	stmt := "SELECT id, email, fullname, image, status, date_added  FROM staff;"
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

// // This will return the 10 most recently created members.
// func Latest(env *handlers.Env) ([]*models.Staff, error) {
// 	return nil, nil
// }

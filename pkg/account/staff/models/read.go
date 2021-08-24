package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Staff struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullname"`
	Image     string    `json:"image"`
	Status    bool      `json:"status"`
	DateAdded time.Time `json:"dateAdded"`
}

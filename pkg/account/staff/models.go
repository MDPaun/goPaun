package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Staff struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	FullName  string    `json:"fullname"`
	Image     string    `json:"image"`
	IP        string    `json:"ip"`
	Status    bool      `json:"status"`
	DateAdded time.Time `json:"dateAdded"`
}

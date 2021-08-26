package config

import (
	"log"

	"github.com/MDPaun/goPaun/pkg/account/staff/postgres"
)

type Env struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	// DB       *sql.DB
	Staff *postgres.StaffModel
}

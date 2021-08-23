package config

import (
	"database/sql"
	"log"
)

type Env struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	DB       *sql.DB
}

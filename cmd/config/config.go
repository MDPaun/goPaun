package config

import (
	"html/template"
	"log"

	"github.com/MDPaun/goPaun/pkg/account/staff/postgres"
)

type Env struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	Staff         *postgres.StaffModel
	TemplateCache map[string]*template.Template
}

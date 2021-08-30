package config

import (
	"html/template"
	"log"

	"github.com/MDPaun/goPaun/pkg/account/staff/postgres"
	"github.com/MDPaun/goPaun/pkg/store/inventory/mysqlDecoCraft"
)

type Env struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger

	Staff     *postgres.StaffModel
	Inventory *mysqlDecoCraft.InventoryModel

	TemplateCache map[string]*template.Template
}

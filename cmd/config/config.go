package config

import (
	"html/template"
	"log"

	"github.com/MDPaun/goPaun/pkg/account/staff/postgres"
	"github.com/MDPaun/goPaun/pkg/store/inventory/mysqlDecoCraft"
	inventory "github.com/MDPaun/goPaun/pkg/store/inventory/postgres"
)

type Env struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger

	Staff       *postgres.StaffModel
	Inventory   *inventory.InventoryModel
	InventoryDC *mysqlDecoCraft.InventoryModel

	TemplateCache map[string]*template.Template
}

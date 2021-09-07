package config

import (
	"html/template"
	"log"

	"github.com/MDPaun/goPaun/pkg/account/staff/postgres"
	"github.com/MDPaun/goPaun/pkg/store/inventory/mysqlDecoCraft"
	"github.com/MDPaun/goPaun/pkg/store/inventory/mysqlMercerie"
	inventory "github.com/MDPaun/goPaun/pkg/store/inventory/postgres"
)

type Env struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger

	Staff     *postgres.StaffModel
	Inventory *inventory.InventoryModel
	// FilterProducts *fp.FilterProducts
	InventoryDC *mysqlDecoCraft.InventoryModel
	InventoryMC *mysqlMercerie.InventoryModel

	TemplateCache map[string]*template.Template
}

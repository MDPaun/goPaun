package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	// base "github.com/MDPaun/goPaun/cmd/base"
	"github.com/MDPaun/goPaun/cmd/config"
	"github.com/MDPaun/goPaun/pkg/account/staff/postgres"
	storage "github.com/MDPaun/goPaun/pkg/storage"
	"github.com/MDPaun/goPaun/pkg/store/inventory/mysqlDecoCraft"
	inventory "github.com/MDPaun/goPaun/pkg/store/inventory/postgres"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db := storage.ConnectDB()
	defer db.Close()

	dbDC := storage.ConnectDecoCraft()
	defer dbDC.Close()

	// Initialize a new template cache...
	templateCache, err := config.NewTemplateCache("./../ui/admin/")
	if err != nil {
		errorLog.Fatal(err)
	}

	env := &config.Env{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		Staff:         &postgres.StaffModel{DB: db},
		Inventory:     &inventory.InventoryModel{DB: db},
		InventoryDC:   &mysqlDecoCraft.InventoryModel{DBDC: dbDC},
		TemplateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: env.ErrorLog,
		Handler:  routes(env),
	}
	env.InfoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	env.ErrorLog.Fatal(err)
}

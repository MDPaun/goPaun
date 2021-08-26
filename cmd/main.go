package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	// base "github.com/MDPaun/goPaun/cmd/base"
	"github.com/MDPaun/goPaun/cmd/config"
	"github.com/MDPaun/goPaun/pkg/account/staff/postgres"
	psDB "github.com/MDPaun/goPaun/pkg/storage"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db := psDB.ConnectDB()
	defer db.Close()

	// Initialize a new template cache...
	templateCache, err := config.NewTemplateCache("./../ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	env := &config.Env{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		Staff:         &postgres.StaffModel{DB: db},
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

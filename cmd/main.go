package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	// "github.com/MDPaun/goPaun/pkg/account/staff/models/postgres"
	"github.com/MDPaun/goPaun/cmd/config"
	"github.com/MDPaun/goPaun/pkg/account/staff/models/postgres"
	psDB "github.com/MDPaun/goPaun/pkg/storage"
)

// type application struct {
// 	errorLog *log.Logger
// 	infoLog  *log.Logger
// 	// staff    *postgres.StaffModel
// }

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db := psDB.ConnectDB()
	defer db.Close()

	env := &config.Env{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		// DB:       db,
		Staff: &postgres.StaffModel{DB: db},
	}

	// app := &application{
	// 	errorLog: errorLog,
	// 	infoLog:  infoLog,
	// staff:    &postgres.StaffModel{DB: db},
	// }

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: env.ErrorLog,
		// Handler:  env.routes(),
		Handler: routes(env),
	}
	env.InfoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	env.ErrorLog.Fatal(err)
}

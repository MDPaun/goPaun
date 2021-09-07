package storage

import (
	//for connecting to db mysql
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// define('DB_DATABASE', 'decocraf_decocraft3');

const (
	hostME     = "37.251.160.44"
	portME     = "3306"
	userME     = "merceriedc"
	passwordME = "*Zet0-L^n4~u"
	dbnameME   = "mercerie_dc"
)

// define('DB_USERNAME', 'merceriedc');
// define('DB_PASSWORD', '*Zet0-L^n4~u');
// define('DB_DATABASE', 'mercerie_dc');

func ConnectMercerie() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		userME, passwordME, hostME, portME, dbnameME)

	var err error
	DbMercerie, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = DbMercerie.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected - Mercerie!")
	return DbMercerie
}

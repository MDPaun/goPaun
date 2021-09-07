package storage

import (
	//for connecting to db mysql
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// define('DB_DATABASE', 'decocraf_decocraft3');

const (
	hostDC     = "94.24.54.44"
	portDC     = "3306"
	userDC     = "decocraf_admin2"
	passwordDC = "7y+s$+Bs1VJh"
	dbnameDC   = "decocraf_decocraft3.2"
)

func ConnectDecoCraft() *sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		userDC, passwordDC, hostDC, portDC, dbnameDC)

	var err error
	DbDecoCraft, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	err = DbDecoCraft.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected! - DecoCraft")
	return DbDecoCraft
}

package storage

import (
	"database/sql"
	"fmt"

	//for connecting to db
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432 //5050
	user     = "root"
	password = "root"
	dbname   = "postgres"
)

func ConnectDB() *sql.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DbClient, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = DbClient.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return DbClient
}

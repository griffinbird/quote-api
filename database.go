package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	POSTGRES_DRIVER ="postgres"
	POSTGRES_LOCATION ="postgres://postgres:postgres@10.0.75.1:5432/postgres?sslmode=disable"
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func ExecDB(sqlStatement string, args ...interface{}) (sql.Result, error) {
	db := dbconnect()
	defer db.Close()

	result, err := db.Exec(sqlStatement, args...)
	return result, err
}

func QueryDB(sqlStatement string, args ...interface{}) (*sql.Rows, error) {
	db := dbconnect()
	defer db.Close()

	rows, err := db.Query(sqlStatement, args...)
	return rows, err
}

/*
func dbconnect() *sql.DB {
	db, err := sql.Open(POSTGRES_DRIVER, POSTGRES_LOCATION)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
*/

func dbconnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//fmt.Println("About to open the DB connection")
	db, err := sql.Open(POSTGRES_DRIVER, psqlInfo)
	if err != nil {
	//	fmt.Println(err)
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
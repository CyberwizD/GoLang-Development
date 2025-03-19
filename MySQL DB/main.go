package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("MySQL DB")

	// Database connection.``
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Perform a db.Query insert
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// if there is an error inserting...
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

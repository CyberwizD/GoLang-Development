package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fmt.Println("MySQL DB")

	// Database connection.``
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Perform a db.Query INSERT
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

	// if there is an error inserting...
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	// Perform a db.Query SELECT
	results, err := db.Query("SELECT id, name FROM test")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var tag Tag

		// For each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Name)

		if err != nil {
			panic(err.Error())
		}

		// Print out the tag's ID and Name attribute
		fmt.Println(tag.ID, tag.Name)
	}
}

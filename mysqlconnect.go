package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type TablesList struct {
	table_name string `json:"Tables_in_mstr_new"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "mstr:mstr@tcp(127.0.0.1:3306)/mstr_new")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// perform a db.Query insert
	cities, err := db.Query("show tables;")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	for cities.Next() {
		var table TablesList
		// for each row, scan the result into our tag composite object
		err = cities.Scan(&table.table_name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(table.table_name)
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

}

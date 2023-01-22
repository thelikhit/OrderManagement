package main

import (
	"database/sql"
	"fmt"
)

// testDBConn is a helper function to check successful connection to MySQL.
// Returns error if connection failed.
func testDBConn() error {

	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@123@tcp(0.0.0.0:3306)/mysql")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("unable to close db connection")
		}
	}(db)

	// Test the connection to the database
	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("successfully connected to database")
	return nil
}

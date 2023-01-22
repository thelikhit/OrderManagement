package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {

	// verify successful connection to MySQL
	err := testDBConn()
	if err != nil {
		fmt.Println("unable to connect to database")
		return
	}

	router := Router()
	fmt.Println("starting server...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Order struct {
	ID           string  `json:"id"`
	Status       string  `json:"status"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total"`
	CurrencyUnit string  `json:"currencyUnit"`
}

type Item struct {
	ID          string  `json:"id"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

// hopePage is the handler function for the default endpoint "/".
func homePage(http.ResponseWriter, *http.Request) {
	fmt.Println("SellerApp Assignment - Order Management Service")
}

// getOrders is the handler function for "/orders" endpoint.
// Returns all the orders from the database.
func getOrders(w http.ResponseWriter, _ *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@123@tcp(0.0.0.0:3306)/mysql")
	if err != nil {
		fmt.Println(err)
	}
	// Close connection after function ends to avoid memory leak
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	// fetch data from Orders table
	var orders []Order
	rows, err := db.Query("SELECT * FROM Orders")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// iterate over the returned rows
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.Status, &order.Total, &order.CurrencyUnit)
		orderID := order.ID
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Fetch data from Items table based on orderID
		var items []Item
		query := "SELECT ItemID, description, price, quantity FROM Items WHERE orderID = " + "\"" + orderID + "\"" + ";"
		rows2, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// iterate over the returned rows
		for rows2.Next() {
			var item Item
			if err = rows2.Scan(&item.ID, &item.Description, &item.Price, &item.Quantity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			items = append(items, item)
		}
		order.Items = items
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode result into JSON
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getOrdersByField is the handler function for "/orders/{field}" endpoint.
// Returns all orders from the database sorted by 'field'.
func getOrdersByField(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@123@tcp(0.0.0.0:3306)/mysql")
	if err != nil {
		fmt.Println(err)
	}
	// Close connection after function ends to avoid memory leak
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	// params is a map that contains the path parameters passed to the endpoint
	params := mux.Vars(r)

	// fetch data from Orders table
	var orders []Order
	query := "SELECT * FROM Orders ORDER BY " + params["field"] + ";"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// iterate over the returned rows
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.Status, &order.Total, &order.CurrencyUnit)
		orderID := order.ID
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Fetch data from Items table based on orderID
		var items []Item
		query := "SELECT ItemID, description, price, quantity FROM Items WHERE orderID = " + "\"" + orderID + "\"" + ";"
		rows2, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// iterate over the returned rows
		for rows2.Next() {
			var item Item
			if err = rows2.Scan(&item.ID, &item.Description, &item.Price, &item.Quantity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			items = append(items, item)
		}
		order.Items = items
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode into JSON
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getOrder is the handler function for "/order/{key}/{value}" endpoint.
// Returns order(s) matching the key-value pair - key:value.
func getOrder(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@123@tcp(0.0.0.0:3306)/mysql")
	if err != nil {
		fmt.Println(err)
	}
	// Close connection after function ends to avoid memory leak
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	params := mux.Vars(r)

	// fetch data from Orders table
	var orders []Order
	query := "SELECT * FROM Orders WHERE " + params["key"] + " = " + "\"" + params["value"] + "\"" + ";"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.ID, &order.Status, &order.Total, &order.CurrencyUnit)
		orderID := order.ID
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Fetch data from Items table based on orderID
		var items []Item
		query := "SELECT ItemID, description, price, quantity FROM Items WHERE orderID = " + "\"" + orderID + "\"" + ";"
		rows2, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for rows2.Next() {
			var item Item
			if err = rows2.Scan(&item.ID, &item.Description, &item.Price, &item.Quantity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			items = append(items, item)
		}
		order.Items = items
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode into JSON
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// addOrder is the handler function for "/add". Adds the JSON passed as body to the database.
// Returns a status code.
func addOrder(w http.ResponseWriter, r *http.Request) {

	// set headers
	w.Header().Set("Content-Type", "application/json")

	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@123@tcp(0.0.0.0:3306)/mysql")
	if err != nil {
		panic(err.Error())
	}
	// Close connection after function ends to avoid memory leak
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("error in closing db connection")
			panic(err.Error())
		}
	}(db)

	// Parse the JSON payload from the request body
	var order Order
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert the data into the MySQL database
	query := "INSERT INTO orders (orderID, status, total, currencyUnit) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, order.ID, order.Status, order.Total, order.CurrencyUnit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Order inserted into the database")

	// Inserting the items in the separate table
	for _, item := range order.Items {
		query = "INSERT INTO items (itemID, orderID, description, price, quantity) VALUES (?, ?, ?, ?, ?)"
		_, err := db.Exec(query, item.ID, order.ID, item.Description, item.Price, item.Quantity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	fmt.Println("data inserted into database")

}

// updateOrder is the handler function for "/update". Updates the status of an existing order.
// Returns a status code.
func updateOrderStatus(w http.ResponseWriter, r *http.Request) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	// Open a connection to the database
	db, err := sql.Open("mysql", "root:password@123@tcp(0.0.0.0:3306)/mysql")
	if err != nil {
		panic(err.Error())
	}
	// Close connection after function ends to avoid memory leak
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("error in closing db connection")
		}
	}(db)

	// Parse the JSON payload from the request body
	var order Order
	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update status of order
	query := "UPDATE orders SET status = ? WHERE orderID = ?"
	_, err = db.Exec(query, order.Status, order.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Order updated into the database")
}

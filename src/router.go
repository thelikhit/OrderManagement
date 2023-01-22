package main

import (
	"github.com/gorilla/mux"
)

// Router handles all the requests made to the endpoint for this service.
func Router() *mux.Router {

	// router instance
	router := mux.NewRouter().StrictSlash(true)

	// endpoint for homepage
	router.HandleFunc("/", homePage).Methods("GET")

	// endpoint to get all orders
	router.HandleFunc("/orders", getOrders).Methods("GET")

	// endpoint to get orders sorted by field
	router.HandleFunc("/orders/{field}", getOrdersByField).Methods("GET")

	//endpoint to get a particular order
	router.HandleFunc("/order/{key}/{value}", getOrder).Methods("GET")

	// endpoint to add a new order
	router.HandleFunc("/add", addOrder).Methods("POST")

	// endpoint to update the status of an existing order
	router.HandleFunc("/update", updateOrderStatus).Methods("PUT")

	return router
}

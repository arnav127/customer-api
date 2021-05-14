package main

import (
	"fmt"
	"net/http"
)

func main() {
	//make map to store user details

	//Api endpoints
	http.HandleFunc("/users/create", createCustomer)
	http.HandleFunc("/users/update", updateCustomer)
	http.HandleFunc("/users/delete", deleteCustomer)
	http.HandleFunc("/users/list", listAllCustomer)
	http.HandleFunc("/users/search", searchCustomer)

	fmt.Println("Server started")

	//Starting server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
		return
	}
}


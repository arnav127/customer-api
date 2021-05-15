package main

import (
	"fmt"
	"net/http"
)

func main() {
	//make map to store user details

	//Api endpoints
	http.HandleFunc("/v1/users/create", createCustomer)
	http.HandleFunc("/v1/users/update", updateCustomer)
	http.HandleFunc("/v1/users/delete", deleteCustomer)
	http.HandleFunc("/v1/users/list", listAllCustomer)
	http.HandleFunc("/v1/users/search", searchCustomer)

	fmt.Println("Server started")

	//Starting server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
		return
	}
}

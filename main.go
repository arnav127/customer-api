package main

import (
	"fmt"
	"net/http"
)

func main() {
	//make map to store user details

	//v1 api endpoints
	http.HandleFunc("/v1/users/create", createCustomer)
	http.HandleFunc("/v1/users/update", updateCustomer)
	http.HandleFunc("/v1/users/delete", deleteCustomer)
	http.HandleFunc("/v1/users/list", listAllCustomer)
	http.HandleFunc("/v1/users/search", searchCustomer)

	//v2 api endpoints
	http.HandleFunc("/v2/user", handleCustomer)
	http.HandleFunc("/v2/users/", listAllCustomer)

	fmt.Println("Server started")

	//Starting server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
		return
	}
}

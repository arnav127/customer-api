package main

import (
	"fmt"
	"net/http"
)

func deleteCustomer(response http.ResponseWriter, request *http.Request) {
	//TODO: remove customer record
	fmt.Println("Delete")
}

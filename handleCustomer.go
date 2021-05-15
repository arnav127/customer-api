package main

import "net/http"

func handleCustomer(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		createCustomer(response, request)
	case "GET":
		searchCustomer(response, request)
	case "DELETE":
		deleteCustomer(response, request)
	case "PUT":
		updateCustomer(response, request)
	case "PATCH":
		updateCustomer(response, request)
	default:
		response.Header().Set("content-type", "application/json")
		response.WriteHeader(http.StatusNotImplemented)
		if _, err := response.Write([]byte(`{ "error": "Method not implemented" }`)); err != nil {
			panic(err)
		}
		return
	}
}

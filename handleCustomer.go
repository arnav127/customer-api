package main

import "net/http"

func handleCustomer(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		createCustomer(response, request)
	case http.MethodGet:
		searchCustomer(response, request)
	case http.MethodDelete:
		deleteCustomer(response, request)
	case http.MethodPut:
		updateCustomer(response, request)
	case http.MethodPatch:
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

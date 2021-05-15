package main

import (
	"net/http"
)

func deleteCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	//Only GET and DELETE requests allowed
	if request.Method != http.MethodDelete {
		response.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := response.Write([]byte(`{ "error": "Method not allowed" }`)); err != nil {
			panic(err)
		}
		return
	}

	// Check if id present in URL
	deleteID := request.URL.Query()["id"]
	if deleteID != nil {
		for idx, user := range users {
			if user.Id == deleteID[0] {
				//Swap user to delete with last user in slice and reduce slice length by 1
				totalUsers := len(users)
				users[totalUsers-1], users[idx] = users[idx], users[totalUsers-1]
				users = users[:totalUsers-1]
				response.WriteHeader(http.StatusNoContent)
				return
			}
		}
	} else {
		//Id not present in parameter
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error": "Bad Request" }`)); err != nil {
			panic(err)
		}
	}

	//return error when user not found
	response.WriteHeader(http.StatusNotFound)
	if _, err := response.Write([]byte(`{ "error": "Customer does not exist" }`)); err != nil {
		panic(err)
	}

}

package main

import (
	"fmt"
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
		var user User
		queryString := fmt.Sprintf("delete from users where id='%v' returning *", deleteID[0])
		query, err := db.Query(queryString)
		if err != nil {
			panic(err)
		}
		if query.Next() {
			err = query.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
			if err != nil {
				panic(err)
			}
			response.WriteHeader(http.StatusNoContent)
			return
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

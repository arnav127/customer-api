package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func searchCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	//Only GET requests allowed
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := response.Write([]byte(`{ "error": "Method not allowed" }`)); err != nil {
			panic(err)
		}
		return
	}

	//Query email and first_name from URL
	searchByEmail := request.URL.Query()["email"]
	searchByFirstName := request.URL.Query()["first_name"]
	if searchByEmail != nil && searchByFirstName != nil {
		queryString := fmt.Sprintf("select id, firstname, lastname, email, phone from users where email='%s' and firstname='%s'",
			searchByEmail[0], searchByFirstName[0])
		query, err := db.Query(queryString)
		if err != nil {
			panic(err)
		}
		defer query.Close()
		var userResult User
		if query.Next() {
			err = query.Scan(&userResult.Id, &userResult.FirstName, &userResult.LastName, &userResult.Email, &userResult.Phone)
			if err != nil {
				panic(err)
			}
			if err := json.NewEncoder(response).Encode(userResult); err != nil {
				panic(err)
			}
			return
		}
	} else {
		//Both email and first_name cannot be queried from URL
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{"error": "Bad Request"}`)); err != nil {
			panic(err)
		}
		return
	}

	//Return NotFound if user not in slice
	response.WriteHeader(http.StatusNotFound)
	if _, err := response.Write([]byte(`{"error": "Record not found"}`)); err != nil {
		panic(err)
	}
}

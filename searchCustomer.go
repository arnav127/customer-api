package main

import (
	"database/sql"
	"encoding/json"
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

		var userResult User
		const queryString = "select id, firstname, lastname, email, phone from users where email=$1 and firstname=$2"

		//Query database and generate userResult from row returned
		row := db.QueryRow(queryString, searchByEmail[0], searchByFirstName[0])
		err := row.Scan(&userResult.Id, &userResult.FirstName, &userResult.LastName, &userResult.Email, &userResult.Phone)

		switch err {
		//row returned successfully
		case nil:
			if err := json.NewEncoder(response).Encode(userResult); err != nil {
				panic(err)
			}
		//no row returned
		case sql.ErrNoRows:
			response.WriteHeader(http.StatusNotFound)
			if _, err := response.Write([]byte(`{"error": "Record not found"}`)); err != nil {
				panic(err)
			}
		default:
			panic(err)
		}
	} else {
		//Both email and first_name cannot be queried from URL
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{"error": "Bad Request"}`)); err != nil {
			panic(err)
		}
	}
}

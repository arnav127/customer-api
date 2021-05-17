package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
)

func getCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	//Only GET requests allowed
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := response.Write([]byte(`{ "error": "Method not allowed" }`)); err != nil {
			panic(err)
		}
		return
	}
	id := path.Base(request.URL.Path)

	var userResult User
	const queryString = "select id, firstname, lastname, email, phone from users where id=$1"
	//Query database and generate userResult from row returned
	row := db.QueryRow(queryString, id)
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
}

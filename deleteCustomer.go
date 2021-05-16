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
		//var user User
		const queryString = "delete from users where id=$1 returning *"
		query, err := db.Exec(queryString, deleteID[0])
		if err != nil {
			panic(err)
		}
		rowsAffected, err := query.RowsAffected()
		if err != nil {
			if _, err := response.Write([]byte(`{ "error": "` + err.Error() +`" }`)); err != nil {
				panic(err)
			}
		}
		switch rowsAffected{
		case 1:
			response.WriteHeader(http.StatusNoContent)
		default:
			response.WriteHeader(http.StatusNotFound)
			if _, err := response.Write([]byte(`{ "error": "Customer does not exist" }`)); err != nil {
				panic(err)
			}
		}

	} else {
		//Id not present in parameter
		response.WriteHeader(http.StatusBadRequest)
		if _, err := response.Write([]byte(`{ "error": "Bad Request" }`)); err != nil {
			panic(err)
		}
	}
}

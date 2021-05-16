package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {

	//connect to the database
	connStr := genConnectionString()
	fmt.Println(connStr)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

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

func genConnectionString() string {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("SSLMODE")
	return fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=%v", user, password, dbname, port, sslmode)
}

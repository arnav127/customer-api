package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/arnavdixit/customer-api/server"
	"google.golang.org/grpc"
	"net"
	"os"
)

//var db *sql.DB

func main() {

	//connect to the database
	connStr := genConnectionString()
	fmt.Println(connStr)
	var err error
	server.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := server.Server{}
	grpcServer := grpc.NewServer()

	server.RegisterCustomerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
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

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/arnavdixit/customer-api/core"
	"gitlab.com/arnavdixit/customer-api/proto"
	"google.golang.org/grpc"
	"net"
	"os"
)

func initService () *sql.DB{

	connStr := genConnectionString()
	fmt.Println(connStr)
	var err error
	Db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return Db

}


func genConnectionString() string {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := os.Getenv("PORT")
	sslmode := os.Getenv("SSLMODE")
	return fmt.Sprintf("user=%v password=%v dbname=%v port=%v sslmode=%v", user, password, dbname, port, sslmode)
}

//var db *sql.DB

func main() {

	//connect to the database

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	db := initService()
	cs := core.Service{
		Db: db,
	}

	//create user schema
	cs.Initdb()

	s := CustomerServiceController{
		CustomerService: &cs,
	}
	grpcServer := grpc.NewServer()

	proto.RegisterCustomerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"gitlab.com/arnavdixit/customer-api/proto"
	"google.golang.org/grpc"
)

var options = []string{
	"Create New Customer",
	"Get Customer by Id",
	"Get All Customers",
	"Search Customer by Email and Name",
	"Update Customer Record",
	"Delete Existing Customer",
	"Exit",
}

var client proto.CustomerServiceClient

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client = proto.NewCustomerServiceClient(conn)
	fmt.Println("Connected to the server!")
	quit := false
	for !quit {
		switch displayMenu() {
		case 1:
			create()
		case 2:
			getbyid()
		case 3:
			getall()
		case 4:
			search()
		case 5:
			update()
		case 6:
			deletecus()
		case 7:
			quit = true
		default:
			fmt.Println("Please enter correct option!")
		}
		fmt.Println("-------------------------------")
	}
	fmt.Println("")

}

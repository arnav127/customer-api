package main

import (
	"context"
	"fmt"
	"gitlab.com/arnavdixit/customer-api/server"
	"google.golang.org/grpc"
	"io"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := server.NewCustomerServiceClient(conn)
	getreq := server.GetCustomerRequest{Id: "14"}
	resp, err := c.GetCustomer(context.Background(), &getreq)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
	var u *server.User
	r2, err := c.GetAllCustomers(context.Background(), &server.NoQuery{})
	if err != nil {
		panic(err)
	}
	for  {
		u, err = r2.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(u)
	}

}

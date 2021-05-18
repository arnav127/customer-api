package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.com/arnavdixit/customer-api/proto"
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

	c := proto.NewCustomerServiceClient(conn)
	getreq := proto.GetCustomerRequest{Id: "14"}
	resp, err := c.GetCustomer(context.Background(), &getreq)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
	var u *proto.User
	r2, err := c.GetAllCustomers(context.Background(), &empty.Empty{})
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

package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.com/arnavdixit/customer-api/proto"
	"os"
	"text/tabwriter"
)

func create() {
	var user proto.User
	fmt.Print("Id: ")
	fmt.Scanf("%v", &user.Id)
	fmt.Print("First Name: ")
	fmt.Scanf("%v", &user.FirstName)
	fmt.Print("Last Name: ")
	fmt.Scanf("%v", &user.LastName)
	fmt.Print("Email: ")
	fmt.Scanf("%v", &user.Email)
	fmt.Print("Phone: ")
	fmt.Scanf("%v", &user.Phone)
	var createUser proto.CreateUserRequest = proto.CreateUserRequest{User: &user}

	resp, err := client.CreateCustomerService(context.Background(), &createUser)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User created: ")
	fmt.Println(Teal("Id: "), Yellow(resp.Id))
	fmt.Println(Teal("First Name: "), Yellow(resp.FirstName))
	fmt.Println(Teal("Last Name: "), Yellow(resp.LastName))
	fmt.Println(Teal("Email: "), Yellow(resp.Email))
	fmt.Println(Teal("Phone: "), Yellow(resp.Phone))
}

func getall() {
	var u *proto.User

	r2, err := client.GetAllCustomers(context.Background(), &empty.Empty{})

	if err != nil {
		panic(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 5, 15, 0, ' ', tabwriter.Debug)
	fmt.Println("Customers in database: ")
	fmt.Fprintln(w, Teal("Id"), "\t", Teal("First Name"), "\t", Teal("Last Name"), "\t", Teal("Email"), "\t", Teal("Phone"))
	fmt.Fprintln(w, Teal("----"), "\t", Teal("----------"), "\t", Teal("---------"), "\t", Teal("-----"), "\t", Teal("----------"))
	for {
		u, err = r2.Recv()
		if err != nil {
			break
		}
		fmt.Fprintln(w, Yellow(u.Id), "\t", Yellow(u.FirstName), "\t", Yellow(u.LastName), "\t", Yellow(u.Email), "\t", Yellow(u.Phone))
	}
	w.Flush()
}

func getbyid() {
	var getCustomer proto.GetCustomerRequest

	fmt.Print("Enter Customer Id: ")
	fmt.Scanf("%s", &getCustomer.Id)

	resp, err := client.GetCustomer(context.Background(), &getCustomer)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Teal("Id: "), Yellow(resp.Id))
		fmt.Println(Teal("First Name: "), Yellow(resp.FirstName))
		fmt.Println(Teal("Last Name: "), Yellow(resp.LastName))
		fmt.Println(Teal("Email: "), Yellow(resp.Email))
		fmt.Println(Teal("Phone: "), Yellow(resp.Phone))
	}
}

func search() {
	var searchRequest proto.SearchCustomerRequest

	fmt.Print("Email of Customer: ")
	fmt.Scanf("%v", &searchRequest.Email)
	fmt.Print("First Name of Customer: ")
	fmt.Scanf("%v", &searchRequest.FirstName)

	resp, err := client.SearchCustomer(context.Background(), &searchRequest)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Teal("Id: "), Yellow(resp.Id))
		fmt.Println(Teal("First Name: "), Yellow(resp.FirstName))
		fmt.Println(Teal("Last Name: "), Yellow(resp.LastName))
		fmt.Println(Teal("Email: "), Yellow(resp.Email))
		fmt.Println(Teal("Phone: "), Yellow(resp.Phone))
	}
}

func deletecus() {
	var deleteRequest proto.DeleteCustomerRequest

	fmt.Print("Enter Customer Id: ")
	fmt.Scanf("%s", &deleteRequest.Id)

	_, err := client.DeleteCustomer(context.Background(), &deleteRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User deleted!")
	}
}

func update() {
	var user proto.User
	var id string
	fmt.Print("Id of Customer to Update: ")
	fmt.Scanf("%v", &id)
	fmt.Println("Enter Updated Details(to skip field press enter):")
	fmt.Print("Updated Id: ")
	fmt.Scanf("%v", &user.Id)
	fmt.Print("Updated First Name: ")
	fmt.Scanf("%v", &user.FirstName)
	fmt.Print("Updated Last Name: ")
	fmt.Scanf("%v", &user.LastName)
	fmt.Print("Updated Email: ")
	fmt.Scanf("%v", &user.Email)
	fmt.Print("Updated Phone: ")
	fmt.Scanf("%v", &user.Phone)

	var updateRequest proto.UpdateCustomerRequest = proto.UpdateCustomerRequest{
		UserId: id,
		User:   &user,
	}
	resp, err := client.UpdateCustomer(context.Background(), &updateRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated User Details: ")
		fmt.Println(Teal("Id: "), Yellow(resp.Id))
		fmt.Println(Teal("First Name: "), Yellow(resp.FirstName))
		fmt.Println(Teal("Last Name: "), Yellow(resp.LastName))
		fmt.Println(Teal("Email: "), Yellow(resp.Email))
		fmt.Println(Teal("Phone: "), Yellow(resp.Phone))
	}
}

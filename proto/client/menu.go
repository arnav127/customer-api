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
	fmt.Print("\u001B[1;36mId: \u001B[1;33m")
	fmt.Scanf("%v", &user.Id)
	fmt.Print("\u001B[1;36mFirst Name: \u001B[1;33m")
	fmt.Scanf("%v", &user.FirstName)
	fmt.Print("\u001B[1;36mLast Name: \u001B[1;33m")
	fmt.Scanf("%v", &user.LastName)
	fmt.Print("\u001B[1;36mEmail: \u001B[1;33m")
	fmt.Scanf("%v", &user.Email)
	fmt.Print("\u001B[1;36mPhone: \u001B[1;33m")
	fmt.Scanf("%v", &user.Phone)
	fmt.Print("\u001B[0m")
	var createUser proto.CreateUserRequest = proto.CreateUserRequest{User: &user}

	resp, err := client.CreateCustomerService(context.Background(), &createUser)

	if err != nil {
		fmt.Println(Red(err))
	}
	fmt.Print(Purple("User created:\n\n"))
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
	fmt.Print(Purple("Customers in database:\n\n"))
	fmt.Fprintln(w, Teal("Id"), "\t", Teal("First Name"), "\t", Teal("Last Name"), "\t", Teal("Email"), "\t", Teal("Phone"))
	fmt.Fprintln(w, Teal("--"), "\t", Teal("----------"), "\t", Teal("---------"), "\t", Teal("-----"), "\t", Teal("-----"))
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

	fmt.Print("\u001B[1;36mEnter Customer Id: \u001B[1;33m")
	fmt.Scanf("%s", &getCustomer.Id)
	fmt.Print("\u001B[0m\n")

	resp, err := client.GetCustomer(context.Background(), &getCustomer)

	if err != nil {
		fmt.Println(Red(err))
	} else {
		fmt.Println(Green("Customer found!"))
		fmt.Println(Teal("Id: "), Yellow(resp.Id))
		fmt.Println(Teal("First Name: "), Yellow(resp.FirstName))
		fmt.Println(Teal("Last Name: "), Yellow(resp.LastName))
		fmt.Println(Teal("Email: "), Yellow(resp.Email))
		fmt.Println(Teal("Phone: "), Yellow(resp.Phone))
	}
}

func search() {
	var searchRequest proto.SearchCustomerRequest

	fmt.Print("\u001B[1;36mEmail of Customer: \u001B[1;33m")
	fmt.Scanf("%v", &searchRequest.Email)
	fmt.Print("\u001B[1;36mFirst Name of Customer: \u001B[1;33m")
	fmt.Scanf("%v", &searchRequest.FirstName)
	fmt.Print("\u001B[0m\n")

	resp, err := client.SearchCustomer(context.Background(), &searchRequest)

	if err != nil {
		fmt.Println(Red(err))
	} else {
		fmt.Println(Green("Customer found!"))
		fmt.Println(Teal("Id: "), Yellow(resp.Id))
		fmt.Println(Teal("First Name: "), Yellow(resp.FirstName))
		fmt.Println(Teal("Last Name: "), Yellow(resp.LastName))
		fmt.Println(Teal("Email: "), Yellow(resp.Email))
		fmt.Println(Teal("Phone: "), Yellow(resp.Phone))
	}
}

func deletecus() {
	var deleteRequest proto.DeleteCustomerRequest

	fmt.Print("\u001B[1;36mEnter Customer Id: \u001B[1;33m")
	fmt.Scanf("%s", &deleteRequest.Id)
	fmt.Print("\u001B[0m\n\n")

	_, err := client.DeleteCustomer(context.Background(), &deleteRequest)
	if err != nil {
		fmt.Println(Red(err))
	} else {
		fmt.Println(Green("User deleted!"))
	}
}

func update() {
	var user proto.User
	var id string
	fmt.Print("\u001B[1;36mId of Customer to Update: \u001B[1;33m")
	fmt.Scanf("%v", &id)
	fmt.Println(Purple("Enter Updated Details(to skip field press enter):"))
	fmt.Print("\u001B[1;36mUpdated Id: ")
	fmt.Scanf("%v", &user.Id)
	fmt.Print("\u001B[1;36mUpdated First Name: \u001B[1;33m")
	fmt.Scanf("%v", &user.FirstName)
	fmt.Print("\u001B[1;36mUpdated Last Name: \u001B[1;33m")
	fmt.Scanf("%v", &user.LastName)
	fmt.Print("\u001B[1;36mUpdated Email: \u001B[1;33m")
	fmt.Scanf("%v", &user.Email)
	fmt.Print("\u001B[1;36mUpdated Phone: \u001B[1;33m")
	fmt.Scanf("%v", &user.Phone)
	fmt.Print("\u001B[0m\n\n")

	var updateRequest proto.UpdateCustomerRequest = proto.UpdateCustomerRequest{
		UserId: id,
		User:   &user,
	}
	resp, err := client.UpdateCustomer(context.Background(), &updateRequest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(Green("Updated User Details:\n\n"))
		fmt.Println(Teal("Id: "), Yellow(resp.Id))
		fmt.Println(Teal("First Name: "), Yellow(resp.FirstName))
		fmt.Println(Teal("Last Name: "), Yellow(resp.LastName))
		fmt.Println(Teal("Email: "), Yellow(resp.Email))
		fmt.Println(Teal("Phone: "), Yellow(resp.Phone))
	}
}

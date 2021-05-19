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
			fmt.Println(Red("Please enter correct option!"))
		}
		//fmt.Println(Yellow("-------------------------------"))
	}
	fmt.Println("")

}


func displayMenu() (opt int) {
	fmt.Println("\n***************************")
	for id, option := range options {
		fmt.Println(Purple(id+1) + ") " + Green(option))
	}
	fmt.Print(Magenta("Choose appropriate option: "))
	opt = 0
	fmt.Print("\u001B[1;34m")
	fmt.Scanf("%d", &opt)
	fmt.Print("\u001B[0m\n\n")
	return
}

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}


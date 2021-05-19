package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	customer_api "gitlab.com/arnavdixit/customer-api"
	"gitlab.com/arnavdixit/customer-api/proto"
)

type CustomerServiceController struct {
	proto.UnimplementedCustomerServiceServer
	CustomerService customer_api.Service
}

func genProtoFromUser(user *customer_api.DbUser) *proto.User {
	protoUser := proto.User{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}
	return &protoUser
}

func genUserFromProto(request *proto.User) *customer_api.DbUser {
	createUser := customer_api.DbUser{
		Id:        request.Id,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}
	return &createUser
}

func (ctlr *CustomerServiceController) CreateCustomerService (ctx context.Context, request *proto.CreateUserRequest) (*proto.User, error) {
	fmt.Print("CreateCustomerService: ")
	fmt.Println(request.GetUser())
	err := request.ValidateAll()
	if err != nil {
		return nil, err
	}
	createUser := genUserFromProto(request.User)
	user, err := ctlr.CustomerService.CreateCustomer(createUser)
	if err != nil {
		return nil, err
	}
	protoUser := genProtoFromUser(user)
	return protoUser, nil
}

func (ctlr *CustomerServiceController) GetCustomer (ctx context.Context, request *proto.GetCustomerRequest) (*proto.User, error) {
	fmt.Print("GetCustomer: ")
	fmt.Println(request.GetId())
	err := request.ValidateAll()
	if err != nil {
		return nil, err
	}
	user, err := ctlr.CustomerService.GetCustomer(request.GetId())
	if err != nil {
		return nil, err
	}
	protoUser := genProtoFromUser(user)
	return protoUser, nil
}

func (ctlr *CustomerServiceController) GetAllCustomers(query *empty.Empty, CustomerServer proto.CustomerService_GetAllCustomersServer) error {
	fmt.Println("GET Alllll!!!!")
	userList := ctlr.CustomerService.ListAllCustomer()
	for _, user := range *userList {
		protoUser := genProtoFromUser(&user)
		CustomerServer.Send(protoUser)
	}
	return nil
}
func (ctlr *CustomerServiceController) SearchCustomer (ctx context.Context, request *proto.SearchCustomerRequest) (*proto.User, error) {
	fmt.Print("SearchCustomer: ")
	err := request.ValidateAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(request.Email, request.FirstName)
	user, err := ctlr.CustomerService.SearchCustomer(request.GetEmail(), request.GetFirstName())
	if err != nil {
		return nil, err
	}
	protoUser := genProtoFromUser(user)
	return protoUser, nil
}

func (ctlr *CustomerServiceController) UpdateCustomer (ctx context.Context, request *proto.UpdateCustomerRequest) (*proto.User, error) {
	fmt.Print("UpdateCustomer: ")
	err := request.ValidateAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(request.UserId, request.User)
	dbuser := genUserFromProto(request.User)
	user, err := ctlr.CustomerService.UpdateCustomer(request.GetUserId(), dbuser)
	if err != nil {
		return nil, err
	}
	protoUser := genProtoFromUser(user)
	return protoUser, nil
}

func (ctlr *CustomerServiceController) DeleteCustomer (ctx context.Context, request *proto.DeleteCustomerRequest) (*empty.Empty, error) {
	fmt.Print("DeleteCustomer: ")
	err := request.ValidateAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(request.Id)
	err = ctlr.CustomerService.DeleteCustomer(&request.Id)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

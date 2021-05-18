package main

import (
	"context"
	"fmt"
	customer_api "gitlab.com/arnavdixit/customer-api"
	"gitlab.com/arnavdixit/customer-api/proto"
)

type CustomerServiceController struct {
	//proto.CustomerServiceServer
	proto.UnimplementedCustomerServiceServer
	CustomerService customer_api.Service
}

//func (ctlr *CustomerServiceController) mustEmbedUnimplementedCustomerServiceServer() {
//	panic("implement me")
//}

//func NewCustomerServiceController (CustomerService customer_api.Service) proto.CustomerServiceServer {
//	return &customerServiceController{
//		CustomerService: CustomerService
//	}
//}

//func (ctrl *customerServiceController)
//func genUserFromProto



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
	user, err := ctlr.CustomerService.GetCustomer(request.GetId())
	if err != nil {
		return nil, err
	}
	protoUser := genProtoFromUser(user)
	return protoUser, nil
}

func (ctlr *CustomerServiceController) GetAllCustomers(query *proto.NoQuery, CustomerServer proto.CustomerService_GetAllCustomersServer) error {
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
	fmt.Println(request.UserId, request.User)
	dbuser := genUserFromProto(request.User)
	user, err := ctlr.CustomerService.UpdateCustomer(request.GetUserId(), dbuser)
	if err != nil {
		return nil, err
	}
	protoUser := genProtoFromUser(user)
	return protoUser, nil
}

func (ctlr *CustomerServiceController) DeleteCustomer (ctx context.Context, request *proto.DeleteCustomerRequest) (*proto.NoQuery, error) {
	fmt.Print("DeleteCustomer: ")
	fmt.Println(request.Id)
	err := ctlr.CustomerService.DeleteCustomer(&request.Id)
	if err != nil {
		return nil, err
	}
	return &proto.NoQuery{}, nil
}

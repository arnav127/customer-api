package server

import (
	"context"
	"fmt"
)
//
type Server struct {
}


func (s *Server) mustEmbedUnimplementedCustomerServiceServer() {
	panic("implement me")
}

func (s *Server) CreateCustomerService (ctx context.Context, request *CreateUserRequest) (*User, error) {
	fmt.Print("CreateCustomerService: ")
	fmt.Println(request.GetUser())
	user, err := CreateCustomer(request.User)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) GetCustomer (ctx context.Context, request *GetCustomerRequest) (*User, error) {
	fmt.Print("GetCustomer: ")
	fmt.Println(request.GetId())
	user, err := GetCustomer(request.GetId())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) GetAllCustomers(query *NoQuery, CustomerServer CustomerService_GetAllCustomersServer) error {
	fmt.Println("GET Alllll!!!!")
	userList := ListAllCustomer()
	for _, user := range *userList {
		CustomerServer.Send(&user)
	}
	return nil
}
func (s *Server) SearchCustomer (ctx context.Context, request *SearchCustomerRequest) (*User, error) {
	fmt.Print("SearchCustomer: ")
	fmt.Println(request.Email, request.FirstName)
	user, err := SearchCustomer(request.GetEmail(), request.GetFirstName())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) UpdateCustomer (ctx context.Context, request *UpdateCustomerRequest) (*User, error) {
	fmt.Print("UpdateCustomer: ")
	fmt.Println(request.UserId, request.User)
	user, err := UpdateCustomer(request.GetUserId(), request.User)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) DeleteCustomer (ctx context.Context, request *DeleteCustomerRequest) (*NoQuery, error) {
	fmt.Print("DeleteCustomer: ")
	fmt.Println(request.Id)
	err := DeleteCustomer(&request.Id)
	if err != nil {
		return nil, err
	}
	return &NoQuery{}, nil
}

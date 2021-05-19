package customer_api


// DbUser Struct to store user details
type DbUser struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     int64    `json:"phone"`
}

type Service interface {
	Initdb()
	CreateCustomer(user *DbUser) (*DbUser, error)
	GetCustomer(id string) (*DbUser, error)
	SearchCustomer(email string, firstName string) (*DbUser, error)
	ListAllCustomer() *[]DbUser
	UpdateCustomer(id string, user *DbUser) (*DbUser, error)
	DeleteCustomer(id *string) error
}
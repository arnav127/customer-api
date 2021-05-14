# Customer API

Basic CRUD Customer API in Golang
Domain: localhost:8080

## Customer Struct Fields:  

| Field    | Type  |  
|----------|-------|  
|ID        | String|  
|FIRST NAME| String|  
|LAST NAME | String|  
|EMAIL     | String|  
|PHONE     |    Int|  

## Functionality Provided
- Create customer
- Update customer
- Delete customer
- List All customers
- Search customer

### Create Customer `/users/create`

### Update Customer  `/users/update`

### Delete Customer  `/users/delete?id={id}`

### List All Customers `/users/list`

### Search Customer `/users/search?email={email}&first_name={first_name}`
# Customer API
Basic CRUD Customer API in Golang.  
View the API documentation here: https://documenter.getpostman.com/view/12544023/TzRVgSRg

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

## Status & Error Codes (Version 2 API)

- ### Creating User: `POST /user`
    - Successful Response: `200`
    - Invalid Request: `400`
    - Invalid User Details: `400`

- ### List All Users: `GET /users`
    - Successful Response: `200`
    - Wrong Method: `405`
     
- ### Search User: `GET /user?email={email}&first_name={first_name}`
    - Successful Response: `200`
    - Invalid Request: `400`
    - User Not Found: `404`

- ### Update User: `PUT /user`   `PATCH /user`
    - Successful Response: `200`
    - Invalid Request: `400`
    - Invalid User Details: `400`
    - User Not Found: `404`

- ### Delete User: `DELETE /user`
    - Successful Response: `204`
    - Invalid Request: `400`
    - User Not Found: `404`


## Version 1 endpoints (deprecated)
### Create Customer `/users/create`

### Update Customer  `/users/update`

### Delete Customer  `/users/delete?id={id}`

### List All Customers `/users/list`

### Search Customer `/users/search?email={email}&first_name={first_name}`
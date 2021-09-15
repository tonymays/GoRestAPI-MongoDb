# GoRestAPI-MongoDb
Go REST API Boilerplate for MongoDb

* Author: Anthony Mays
* Current Version: 0.0.1
* Release Date: September 15, 2021
* MIT License
___
## Getting Started

Follow the instructions below to get the Go RESTful API boilerplate installed, configured and running

### Prerequisites
* An Ubuntu 20.04 or CentOS 7 Operating System
* Access to a MongoDB database
* Knowledge of Postman

### Setup
1. Clone the git repo:
```bash
$ git clone https://github.com/tonymays/GoRestAPI-MongoDb.git rest_api
```

2. Obtain package dependencies
```bash
$ cd rest_api
$ go get "github.com/gorilla/mux"
$ go get "github.com/gorilla/handlers"
$ go get "github.com/gofrs/uuid"
$ go get "github.com/dgrijalva/jwt-go"
$ go get "go.mongodb.org/mongo-driver/mongo"
$ go get "golang.org/x/crypto/bcrypt"
$ go get "github.com/lithammer/shortuuid"
$ go get "go.mongodb.org/mongo-driver/bson/primitive"
```

3. Ensure your go path variables are established correctly for this project

4. Copy the conf_example.json file to conf.json (naming must be exact)
```bash
$ cp conf_example.json conf.json
```

5. Open the conf.json file and set it configurable settings based upon the descriptions given below:
```
MongoUri         - Uri connection string to your MongoDB database
DbName           - the name of the MongoDB database
Secret           - the signing secret phrase for JWT Tokens
HTTPS            - [on|off] on if using HTTPS otherwise off
Cert             - the cert file name
Key              - the key file name
ServerListenPort - the port the server is listening on
RootUserName     - the name of the root user on new server initialization
RootPassword     - the password of the root user
Firstname        - the first name of the root user
Lastname         - the last name of the root user
Address          - the address of the root user
City             - the city of the root user
State            - the state of the root user
Zip              - the zip of the root user
Country          - the country of the root user
Email            - the root user's email address
Phone            - the root user's phone number
```

6. Copy src/app/conf_test_example.json to src/app/conf_test.json and repeat step 5 in conf_test.json if you plan on using or expanding the built in go tests.
```bash
$ cd src/app
$ cp conf_test_example.json conf_test.json
```

7. Compile the package - must be in the project folder and your go variables must set for this project
```bash
$ go build ./src/app
```

8. Run the app
```bash
$ ./app
```

* The first time you run the app the server will initialize itself by creating the root user, role, base permissions etc.
```bash
initializing new server
------------------------------------------------------------
   step 1: creating root user...
   step 2: creating admin role...
   step 3: creating system permissions...
   step 4: assigning permissions to the admin role...
   step 5: assigning root user to Admin role...
------------------------------------------------------------
new server initialization complete
Listening on port :8080
```

## API Reference
This package contains 5 separate routers: auth_router, permission_router, role_permission_router, role_router and a user_router

### <ins>I. Router Responsibilities</ins>
* the auth_router is responsible for session management such as sign-in and logout
* the user_router is responsible for managing package user
* the role_router is responsible for managing custom package roles
* the permission_router is responsible for permission management when new features are added
* the role_permission_router is responsible for permission assignment to individual roles

### <ins>II. API List</ins>
* Auth Router
```
POST /auth   - session login
DELETE /auth - session logout
GET /auth    - session refresh
HEAD /auth   - check session
PUT /auth    - change user password
```

* User Router
```
POST /users                            - create user
GET /users                             - find all active users
GET /users/{id}                        - get user specified by {id}
PATCH /users/{id}                      - update user specified by {id}
PUT /users/{id}                        - activate user specified by {id}
DELETE /users/{id}                     - deactivate user specified by {id}
GET /users/{id}/roles                  - get all roles for the user specified by {id}
GET /users/{id}/service_catalog        - get service catalog for the user specified by {id}
PUT /users/{userId}/roles.{roleId}     - assign role specified by {roleId} to user specified by {userId}
PATCH /users/{userId}/roles.{roleId}   - activate role specified by {roleId} to user specified by {userId}
DELETE /users/{userId}/roles.{roleId}  - deactivate role specified by {roleId} to user specified by {userId}
```

* Roles Router
```
POST /roles        - create role
GET /roles         - get all active roles
GET /roles/{id}    - get role specified by {id}
PATCH /roles/{id}  - update role specified by {id}
PUT /roles/{id}    - activate role specified by {id}
DELETE /roles/{id} - deactivate role specified by {id}
```

* Permissions Router
```
POST /permissions         - create permission
GET /permissions          - get all active permissions
GET /permissions/{id}     - get permission specified by {id}
PATCH /permissions/{id}   - update permission specified by {id}
PUT /permissions/{id}     - activate permission specified by {id}
DELETE /permissions/{id}  - deactivate permission specified by {id}
```

* Role Permssions Router
```
POST /role/{id}/permissions  - set permissions for the role specified by {id}
GET /role/{id}/permissions   - get permissions for the role specified by {id}
```

### <ins>III. Auth Router Endpoints</ins>
#### 1. Session Login
* POST /auth

##### Request
* Headers
```
{
	Content-Type: application/json
}
```

* Body
```
{
	"username": "root",
	"password": "abc123xyz890"
}
```
##### Response
* Headers
```
{
	Auth-Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFudGhvbnkuZC5tYXlzQGdtYWlsLmNvbSIsImV4cCI6MTYzMTczNDcyNSwicmVtb3RlX2FkZHIiOiI0Ny4xODUuMjIzLjUxIiwidXNlcl9pZCI6Ijk4MmIyYTc5LWJlMmEtNDdlMi1hNTcyLTEwNjZlYjVhOTljOCIsInVzZXJuYW1lIjoicm9vdCJ9.bv6BuoJhcroMJvt-bB0NnFsJJ5mf4a1U6h4EnKCSY5Q
	Status: 200 OK
}
```

* Body
```
{
    "user_idid": "982b2a79-be2a-47e2-a572-1066eb5a99c8",
    "username": "root",
    "email": "YOUR USER Email Address",
    "remote_addr": "YOUR LOGIN IP",
    "service_catalog": [
        "Can Add User",
        "Can Edit User",
        "Can Delete User"
    ]
}
```

* Call Notes
```
Every API call in this package, minus this API, will require the Auth-Token as a header.  It will be referred as {Auth-Token} going forward.
```

#### 2. Session Logout
* DELETE /auth

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 3. Refresh Session
* GET /auth

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 4. Check Session
* HEAD /auth

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 5. Change User Password
* PUT /auth

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

* Body
```
{
  "username": "root",
  "password": "abc123xyz890",
  "new_password": "xxkd938dkdjs"
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "username": "root"
}
```

### <ins>IV. User Router Endpoints</ins>
#### 1. Create User
* POST /users

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

* Body
```
{
    "username": "root",
    "password": "abc123xy890",
    "first_name": "Jane",
    "last_name": "Doe",
    "address": "123 ABC Ave",
    "city": "Long Beach",
    "state": "Ca",
    "zip": "12345",
    "country": "United States",
    "email": "jane.doe@yehoo.com",
    "phone": "1234567890"
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "user_id": "5d40a6f2-de58-43c6-9992-7fba3198ba16",
    "username": "root",
    "password": "$2a$10$pjq.YwrZR9OMpPiMCNHJcujcW0vMopBsNeyJdemltCWJADIrSwfQ6",
    "first_name": "Jane",
    "last_name": "Doe",
    "address": "123 ABC Ave",
    "city": "Long Beach",
    "state": "CA",
    "zip": "12345",
    "country": "United States",
    "email": "jane.doe@yahoo.com",
    "phone": "123456789",
    "active": "Yes",
    "created": "2021-09-15 21:18:14.868632889 +0000 UTC",
    "modified": "2021-09-15 21:18:14.868632889 +0000 UTC"
}
```

#### 2. Get All Active Users
* GET /users

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
[
    {
        "user_id": "00cd2469-d597-48ac-a864-14aa80009127",
        "username": "root",
        "password": "$2a$10$mvHYyXaszdA8oxMt1dkd1uVpzdAE4g4Lzlv3MwAxQvfXw3r7bhB0O",
        "first_name": "John",
        "last_name": "Doe",
        "address": "123 ABC Ave",
        "city": "Long Beach",
        "state": "Ca",
        "zip": "12345",
        "country": "United States",
        "email": "john.doe@yahoo.com",
        "phone": "1234567890",
        "active": "Yes",
        "created": "2021-09-15 19:58:01.597137168 +0000 UTC",
        "modified": "2021-09-15 19:58:01.597137168 +0000 UTC"
    },
	{
        "user_id": "00cd2469-d597-48ac-a864-14aa80009127",
        "username": "root",
        "password": "$2a$10$mvHYyXaszdA8oxMt1dkd1uVpzdAE4g4Lzlv3MwAxQvfXw3r7bhB0O",
        "first_name": "Jane",
        "last_name": "Doe",
        "address": "123 ABC Ave",
        "city": "Long Beach",
        "state": "Ca",
        "zip": "12345",
        "country": "United States",
        "email": "jane.doe@yahoo.com",
        "phone": "1234567890",
        "active": "Yes",
        "created": "2021-09-15 19:58:01.597137168 +0000 UTC",
        "modified": "2021-09-15 19:58:01.597137168 +0000 UTC"
    },
]
```

#### 3. Get User specified by {id}
* GET /users/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
	"user_id": "00cd2469-d597-48ac-a864-14aa80009127",
	"username": "root",
	"password": "$2a$10$mvHYyXaszdA8oxMt1dkd1uVpzdAE4g4Lzlv3MwAxQvfXw3r7bhB0O",
	"first_name": "John",
	"last_name": "Doe",
	"address": "123 ABC Ave",
	"city": "Long Beach",
	"state": "Ca",
	"zip": "12345",
	"country": "United States",
	"email": "john.doe@yahoo.com",
	"phone": "1234567890",
	"active": "Yes",
	"created": "2021-09-15 19:58:01.597137168 +0000 UTC",
	"modified": "2021-09-15 19:58:01.597137168 +0000 UTC"
},
```

#### 4. Update User
* PATCH /users/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

* Body
```
{
	"phone": "9087654321"
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "phone": "9087654321",
    "modified": "2021-09-15 21:21:56.331564725 +0000 UTC"
}
```

#### 5. Activate User
* PUT /users/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "active": "Yes",
    "modified": "2021-09-15 21:22:48.384153946 +0000 UTC"
}
```

#### 6. Deactivate User
* DELETE /users/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
	"active": "No",
	"modified": "2021-09-15 21:22:48.384153946 +0000 UTC"
}
```

#### 7. User roles for user specified by {id}
* GET /users/{id}/roles

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
[
    {
        "user_role_id": "b0792e8c-818f-438f-836c-631ae79d4d68",
        "user_id": "00cd2469-d597-48ac-a864-14aa80009127",
        "role_id": "c3f4fca4-6949-4a50-b610-6ddaa07be8f9",
        "user_name": "root",
        "role_name": "Admin",
        "active": "Yes",
        "created": "2021-09-15 19:58:01.597137168 +0000 UTC",
        "modified": "2021-09-15 19:58:01.597137168 +0000 UTC"
    }
]
```

#### 8. Get User Service Catalog (Permissions)
* GET /users/{id}/service_catalog

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
[
    "Can Add User",
    "Can Edit User",
    "Can Delete User"
]
```

#### 9. Assign role specified by {roleId} to user specified by {userId}
* PUT /users/{userId}/roles/{roleId}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "user_role_id": "164ae760-def2-420d-8162-30a9f39292d5",
    "user_id": "00cd2469-d597-48ac-a864-14aa80009127",
    "role_id": "8088d468-824b-425e-97d9-c65221164e62",
    "active": "Yes",
    "created": "2021-09-15 21:26:37.71239779 +0000 UTC",
    "modified": "2021-09-15 21:26:37.71239779 +0000 UTC"
}
```

#### 10. Activate role specified by {roleId} for user specified by {userId}
* PATCH /users/{userId}/roles/{roleId}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "active": "Yes",
    "modified": "2021-09-15 21:27:34.484471524 +0000 UTC"
}
```

#### 11. Deactivate role specified by {roleId} for user specified by {userId}
* DELETE /users/{userId}/roles/{roleId}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "active": "No",
    "modified": "2021-09-15 21:27:34.484471524 +0000 UTC"
}
```

### <ins>V. Role Router Endpoints</ins>
#### 1. Create Role
* POST /roles

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

* Body
```
{
    "name": "User"
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "role_id": "8088d468-824b-425e-97d9-c65221164e62",
    "name": "User",
    "active": "Yes",
    "created": "2021-09-15 21:26:00.64853357 +0000 UTC",
    "modified": "2021-09-15 21:26:00.64853357 +0000 UTC"
}
```

#### 2. Get Active Roles
* GET /roles

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
[
    {
        "role_id": "c3f4fca4-6949-4a50-b610-6ddaa07be8f9",
        "name": "Admin",
        "active": "Yes",
        "created": "2021-09-15 19:58:01.64194887 +0000 UTC",
        "modified": "2021-09-15 19:58:01.64194887 +0000 UTC"
    },
    {
        "role_id": "8088d468-824b-425e-97d9-c65221164e62",
        "name": "User",
        "active": "Yes",
        "created": "2021-09-15 21:26:00.64853357 +0000 UTC",
        "modified": "2021-09-15 21:26:00.64853357 +0000 UTC"
    }
]
```

#### 3. Get role specified by {id}
* GET /roles/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
	"role_id": "c3f4fca4-6949-4a50-b610-6ddaa07be8f9",
	"name": "Admin",
	"active": "Yes",
	"created": "2021-09-15 19:58:01.64194887 +0000 UTC",
	"modified": "2021-09-15 19:58:01.64194887 +0000 UTC"
}
```

#### 4. Update role specified by {id}
* PATCH /roles/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

* Body
```
{
    "name": "Users"
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "name": "Users",
    "modified": "2021-09-15 22:28:02.70611011 +0000 UTC"
}
```

#### 5. Activate role specified by {id}
* PUT /roles/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "active": "Yes",
    "modified": "2021-09-15 22:28:52.130154913 +0000 UTC"
}
```

#### 6. Deactivate role specified by {id}
* DELETE /roles/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

* Body
```
{
    "active": "No",
    "modified": "2021-09-15 22:28:52.130154913 +0000 UTC"
}
```

### <ins>VI. Permission Router Endpoints</ins>
#### 1. Create Permission
* POST /permissions

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 2. Get Active Permissions
* GET /permissions

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 3. Get permission specified by {id}
* GET /permissions/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 4. Update permission specified by {id}
* PATCH /permissions/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 5. Activate permissions specified by {id}
* PUT /permissions/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 6. Deactivate permissions specified by {id}
* DELETE /permissions/{id}

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

### <ins>VII. Role Permissions Router Endpoints</ins>
#### 1. Set Role Permissions for the role specified by {id}
* POST /role/{id}/permissions

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

#### 2. Get All Active Permissions for the role specified by {id}
* GET /role/{id}/permissions

##### Request
* Headers
```
{
	Content-Type: application/json
	Auth-Token: {Auth-Token}
}
```

##### Response
* Headers
```
{
	Status: 200 OK
}
```

## API Go Tests









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
* MongoUri         - Uri connection string to your MongoDB database
* DbName           - the name of the MongoDB database
* Secret           - the signing secret phrase for JWT Tokens
* HTTPS            - [on|off] on if using HTTPS otherwise off
* Cert             - the cert file name
* Key              - the key file name
* ServerListenPort - the port the server is listening on
* RootUserName     - the name of the root user on new server initialization
* RootPassword     - the password of the root user
* Firstname        - the first name of the root user
* Lastname         - the last name of the root user
* Address          - the address of the root user
* City             - the city of the root user
* State            - the state of the root user
* Zip              - the zip of the root user
* Country          - the country of the root user
* Email            - the root user's email address
* Phone            - the root user's phone number
```

6. Copy src/app/conf_test_example.json to src/app/conf_test.json and repeat step 5 in conf_test.json if you plan on using or expanding the built in go tests.
```bash
$ cd src/app
$ cp conf_test_example.json conf_test.json
```

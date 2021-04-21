# JWT_REST_Gin_MySQL
Web service CRUD using Golang with GIN for create REST api, MySQL as database, Viper as environment variable, JWT for secure service, redis to store token and Swaggo for API Documentation.

**Prerequisites**

1. [Go](https://golang.org/)
2. [Gin](github.com/gin-gonic/gin)
3. [Mysql](https://www.mysql.com/downloads/)
4. [Viper](https://github.com/spf13/viper)
5. [SQLMock](https://github.com/DATA-DOG/go-sqlmock)
6. [Assert](https://godoc.org/github.com/stretchr/testify/assert)
7. [BCrypt](https://godoc.org/golang.org/x/crypto/bcrypt)
8. [JWT](https://github.com/dgrijalva/jwt-go)
9. [UUID](https://github.com/segmentio/ksuid)
10. [Redis](https://github.com/gomodule/redigo)
11. [Swaggo](https://github.com/swaggo/swag)

## Getting Started
1. Firstly, we need to get Gin, MySQL, Viper, sqlmock, assert, jwt, ksuid for UUID, and redis library dependencies for install it
```
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
go get github.com/spf13/viper
go get github.com/DATA-DOG/go-sqlmock
go get github.com/stretchr/testify/assert
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/segmentio/ksuid
go get github.com/gomodule/redigo/redis
```
2. Download [Redis for Windows](https://github.com/dmajkic/redis/downloads)
3. After you download Redis, you’ll need to extract the executables and then double-click on the redis-server executable.
4. Import dump.sql to your MySQL and configure your credential in folder resource
![Alt text](asset/configureCredentialDB.PNG?raw=true "Configure your credential DB")
5. Open cmd and type `setx APP_ENVIRONMENT STAGING` for default environment
6. Open cmd in your project directory and type `go test -v` , you should get a response similar to the following:
![Alt text](asset/testing_gin.PNG?raw=true "Response Unit Testing")

7. To run application,open cmd in your project directory and type
```
go run main.go
```

## Sample Payload
1. [Login](asset/login.PNG)
2. [Logout](asset/logout.PNG)
3. [Get User By Id](asset/getUserById.PNG)
4. [Get User Detail By Id](asset/getUserDetailById.PNG)
5. [Get All User](asset/getAllUser.PNG)
6. [Get All User Detail](asset/getAllUserDetail.PNG)
7. [Create User](asset/createUser.PNG)
8. [Create User Detail](asset/createUserDetail.PNG)
9. [Update User](asset/updateUser.PNG)
10. [Update User Detail](asset/updateUserDetail.PNG)
11. [Delete User By Id](asset/deleteUserById.PNG)
12. [Delete User Detail By Id](asset/deleteUserDetailById.PNG)
13. [Example error response,in case Update User Detail](asset/updateUserDetailError.PNG)


## Implement Swaggo Documentation
1. Open project directory which contains the `main.go`. Generate the swagger docs with the `swag init` command that wrap in the bash `.\swaggo.sh`

![Alt text](asset/swag_init_sh.PNG?raw=true "Swagg init")

2. Import __swaggo__ dependencies:
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```
3. After installation, add [General API](https://github.com/swaggo/swag#general-api-info) annotation in `main.go` code, for example
```
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	var err error

	// Setup database
	configuration.DB, err = configuration.SetupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer configuration.DB.Close()

	port := viper.GetString("PORT")

	docs.SwaggerInfo.Title = "Swagger Service API"
	docs.SwaggerInfo.Description = "This is service API documentation."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Setup router
	router := router.NewRoutes()
	url := swgGin.URL("http://localhost:" + port + "/swagger/doc.json")
	router.GET("/swagger/*any", swgGin.WrapHandler(swgFiles.Handler, url))

	log.Fatal(router.Run(":" + port))
}
```
4. Add [API Operation](https://github.com/swaggo/swag#api-operation) annotations in controller/service code
```
// getUserByID godoc
// @Summary show master user by id
// @Description get string by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.MUser
// @Failure 400 {string} string
// @Failure 404 {object} model.MUser
// @Failure 500 {string} string
// @Security bearerAuth
// @Router /user/{id} [get]
func getUserByID(c *gin.Context) {
	var user model.MUser
	paramID := c.Param("id")
	varID, err := strconv.ParseInt(paramID, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err = repository.GetUserByID(varID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if (model.MUser{}) == user {
		c.JSON(http.StatusNotFound, user)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
```
5. Generate swaggo documentation
![Alt text](asset/swag_dep.PNG?raw=true "Swagg init parse dependency")

6. Browse Swagger UI [http://localhost:8999/swagger/index.html](http://localhost:8999/swagger/index.html)
![Alt text](asset/swagger_ui.PNG?raw=true "Swagger UI")

7. Execute login authentication endpoint
![Alt text](asset/swaggo_login.PNG?raw=true "Swagger Login Authentication")

8. If put wrong username or / and password,the API endpoit will give output something similiar to the following:
![Alt text](asset/swaggo_login_error.PNG?raw=true "Swagger Login Authentication Bad Credentials")

9. Enter right username and password,then the API endpoit will give output something similiar to the following:
![Alt text](asset/swaggo_login_success.PNG?raw=true "Swagger Login Success")

10. Copy access_token from login response,and place it to Authorize with add a word __Bearer__. This wil be store the token for the rest API.
![Alt text](asset/swaggo_auth.PNG?raw=true "Swagger Authorizations")

11. [Sample Endpoint Get User By Id](asset/swaggo_user_id.PNG)
12. [Sample Endpoint Get All User](asset/swaggo_user_list.PNG)



### NOTES

For those having this problem when run `swag init`:
![Alt text](asset/swag_error.PNG?raw=true "Swagger Error")

Check __$GOPATH__/bin where swag executable is present.
![Alt text](asset/swag_init.PNG?raw=true "Swagger init success")
1. ## Getting Started

2. git clone https://github.com/shahzadthathal/go-gin-rest-mysql-redis.git


3. You need to get Gin, MySQL, Viper, sqlmock, assert, jwt, ksuid for UUID, and redis library dependencies for install it
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

4. Import dump.sql to your MySQL and configure your credential in folder resource

5. To run application,open cmd in your project directory and type
```
go run main.go
```

## Swaggo Documentation
1. Open project directory which contains the `main.go`. Generate the swagger docs with the `swag init` command that wrap in the bash `.\swaggo.sh`

2. Import __swaggo__ dependencies:
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```
3. Browse Swagger UI [http://localhost:8999/swagger/index.html]

## Admin credentials
```
username: admin
password: admin1234
```
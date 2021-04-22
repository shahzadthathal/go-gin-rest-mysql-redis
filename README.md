1. ## Getting Started

2. git clone https://github.com/shahzadthathal/go-gin-rest-mysql-redis.git

3. cd  go-gin-rest-mysql-redis

4. You need to get Gin, MySQL, Viper, sqlmock, assert, jwt, ksuid for UUID, and redis library dependencies for install it. Open cmd in your project directory and type
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

5. Import dump.sql to your MySQL and configure your credential in folder resource


6. Import __swaggo__ dependencies:
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

7. Run Tests
```
go test -v
```

8. Run application
```
go run main.go
```


9.  Browse Swagger UI [http://localhost:8999/swagger/index.html](http://localhost:8999/swagger/index.html)

## Admin credentials
```
username: admin
password: admin1234
```
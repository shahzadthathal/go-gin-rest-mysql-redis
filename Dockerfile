FROM golang:1.14

LABEL maintainer="brullyz@gmail.com"

WORKDIR /golang-app

COPY go.mod .

COPY go.sum .

RUN go mod download

# copy from your dir to container
COPY . .

RUN go build

CMD ["./JWT_REST_Gin_MySQL"]

# ==== image app golang
# open cmd in your project dir
# docker build -t go-gin .
# docker run -p 8081:8999 go-gin
# docker start container image_id

# ==== image db MySQL
# docker network create testing-db-mysql
# docker container run --name mysqltestdb --network testing-db-mysql -v D:/Project/Docker/MySQL:/var/lib/mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=P@ssw0rd MYSQL_DATABASE=ms_account_dev -d mysql
# docker exec -it mysqltestdb bash
# mysql -u root -p
# use mysql
# select host, user from user;
# alter user 'root'@'%' identified with mysql_native_password by 'your_strong_pass';
# flush privileges;
# import / execute dump sql

# ==== image Redis
# docker run --name my-image-redis -d redis
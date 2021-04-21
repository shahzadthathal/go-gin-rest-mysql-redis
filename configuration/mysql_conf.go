package configuration

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var username, password, hostname, dbname string

func initConfMySQL() {
	username = viper.GetString("DB.USER_NAME")
	password = viper.GetString("DB.PASSWORD")
	hostname = viper.GetString("DB.HOST_NAME")
	dbname = viper.GetString("DB.NAME")
}

// DB database global
var DB *sql.DB

// SetupDB db
func SetupDB() (*sql.DB, error) {
	initConfMySQL()

	db, err := sql.Open("mysql", confMysql(dbname))

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func confMysql(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

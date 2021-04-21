package main

import (
	"JWT_REST_Gin_MySQL/configuration"
	"JWT_REST_Gin_MySQL/docs"
	"JWT_REST_Gin_MySQL/router"
	"JWT_REST_Gin_MySQL/util"
	"log"
	"os"

	"github.com/spf13/viper"
	swgFiles "github.com/swaggo/files"
	swgGin "github.com/swaggo/gin-swagger"
)

func init() {

	os.Setenv("APP_ENVIRONMENT", "STAGING")

	// read config environment
	configuration.ReadConfig()

	util.Pool = util.SetupRedisJWT()

}

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

package main

import (
	"gogincosmos/config"
	"gogincosmos/router"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	log.Println("Main - Init")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	log.Println("Start main program")
	appConfig, err := config.SetupApplicationConfig()
	if err != nil {
		log.Fatal("ERROR: Failed to get right configuration setup for Application!")
	}
	config.AppConfig = *appConfig
	log.Println("Configurations are all set", config.AppConfig)

	/*err := config.SetUpEnvForMongoDBConnections()
	if err != nil {
		log.Fatal("Issue with fetch/set ENV variables from KeyVault!")
	}
	if err := config.NewDBConfig(); err != nil {
		log.Println(err)
		log.Fatal("ERROR: Failed to connect to DB")
	}*/

	router := router.AllRoutes()
	log.Println("Running GIN server in port 8080!")
	router.Run(":8080")
}

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
	err := config.SetUpEnvForMongoDBConnections()
	if err != nil {
		log.Fatal("Issue with fetch/set ENV variables from KeyVault!")
	}
	if err := config.NewDBConfig(); err != nil {
		log.Println(err)
		log.Fatal("ERROR: Failed to connect to DB")
	}

	router := router.AllRoutes()
	log.Println("Running GIN server in port 8080!")
	router.Run(":8080")
}

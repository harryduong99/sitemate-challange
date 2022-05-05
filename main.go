package main

import (
	"github/harryduong99/sitemate/databasedriver"
	"github/harryduong99/sitemate/route"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	databasedriver.Mongo.ConnectDatabase()
}

func main() {
	route.InitRoutes()

}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

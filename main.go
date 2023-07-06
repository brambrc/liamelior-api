package main

import (
	"liamelior-api/Database"
	"liamelior-api/Model"
	"liamelior-api/Router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	Router.ServeApps()
	Router.CronJob()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {

	Database.Connect()
	Database.Database.AutoMigrate(&Model.User{})
	Database.Database.AutoMigrate(&Model.Photo{})
	Database.Database.AutoMigrate(&Model.Milestone{})
	Database.Database.AutoMigrate(&Model.Member{})
	Database.Database.AutoMigrate(&Model.TextContent{})
	Database.Database.AutoMigrate(&Model.Schedule{})
}

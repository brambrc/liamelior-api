package main


import (
	"liamelior-api/Database"
	"liamelior-api/Model"
	"liamelior-api/Router"
	"github.com/joho/godotenv"
	"log"
)


func main(){
	loadEnv()
	loadDatabase()
	Router.ServeApps()
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
}



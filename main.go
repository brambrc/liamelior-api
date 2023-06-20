package main


import (
	"liamelior-api/Database"
	"liamelior-api/Controller"
	"liamelior-api/Model"
	"github.com/joho/godotenv"
	"log"
	"github.com/gin-gonic/gin"
	"fmt"
)


func main(){
	loadEnv()
	loadDatabase()
	serveApps()
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


}


func serveApps() {
	router := gin.Default()

	authRoutes := router.Group("/auth")

	authRoutes.POST("/register", Controller.Register)
	authRoutes.POST("/login", Controller.Login)

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}
package Router

import (
	"fmt"
	"liamelior-api/Controller"
	"liamelior-api/Middleware"

	"github.com/gin-gonic/gin"
)

func ServeApps() {
	router := gin.Default()

	authRoutes := router.Group("/auth")
	{
		AuthRoutes(authRoutes)
	}

	contentManagement := router.Group("/content-management")
	{
		ContentManagementRoutes(contentManagement)
	}

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", Controller.Register)
	router.POST("/login", Controller.Login)
}

func ContentManagementRoutes(router *gin.RouterGroup) {
	router.POST("/photo-landing-page", Middleware.AdminMiddleware(), Controller.ContextPhoto)
	router.POST("/text-content", Middleware.AdminMiddleware(), Controller.TextContentStore)
	router.POST("/text-content-update", Middleware.AdminMiddleware(), Controller.TextContentUpdate)
	router.GET("/get-caraousel-photo", Controller.GetCaraouselPhoto)
	router.GET("/gallery", Controller.GetGallery)
}

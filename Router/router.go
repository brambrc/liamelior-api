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

	registerMember := router.Group("/register")
	{
		RegisterRoutes(registerMember)
	}

	milestone := router.Group("/milestone")
	{
		MilestoneRoute(milestone)
	}

	cronJob := router.Group("/cron-job")
	{
		CronJobRoute(cronJob)
	}

	schedule := router.Group("/schedule")
	{
		ScheduleRoute(schedule)
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
	router.PUT("/text-content-update", Middleware.AdminMiddleware(), Controller.TextContentUpdate)
	router.GET("/get-text-content", Controller.GetTextContent)
	router.GET("/get-caraousel-photo", Controller.GetCaraouselPhoto)
	router.GET("/gallery", Controller.GetGallery)
}

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/member", Controller.RegisterMember)
}

func MilestoneRoute(router *gin.RouterGroup) {
	router.POST("/store", Middleware.AdminMiddleware(), Controller.SaveMilestone)
	router.GET("/get", Controller.GetMilestone)
}

func CronJobRoute(router *gin.RouterGroup) {
	router.GET("/cron-get-show-schedule", Controller.ScrapeHandler)
}

func ScheduleRoute(router *gin.RouterGroup) {
	router.GET("/get-upcoming-and-past-shows", Controller.GetUpcomingAndPastShows)
}

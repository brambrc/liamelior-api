package Controller


import (
	"liamelior-api/Model"
	"net/http"
	"github.com/gin-gonic/gin"
)


func GetCaraouselPhoto(context *gin.Context) {

	photo, err := Model.FindPhotoByContext("caraousel")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get photo", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Caraousel retrieved successfully!", "photo": photo})

}
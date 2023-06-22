package Controller

import (
	"liamelior-api/Model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCaraouselPhoto(context *gin.Context) {

	photo, err := Model.FindPhotosByContext("caraousel")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get photo", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Caraousel retrieved successfully!", "photo": photo})

}

func GetGallery(context *gin.Context) {

	//check if request had a query param
	//if it did, get the query param
	param := context.Query("limit")

	if param != "" {
		//convert param into int
		limit, err := strconv.Atoi(param)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get photo", "error": err.Error()})
			return
		}
		photo, err := Model.FindPhotosByContextWithParam("gallery", limit)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get photo", "error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Gallery retrieved successfully!", "photo": photo})

	} else {
		photo, err := Model.FindPhotosByContext("gallery")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get photo", "error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Gallery retrieved successfully!", "photo": photo})

	}

}

package Controller

import (
	"liamelior-api/Model"
	"math"
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
	param := context.Query("limit")
	pageParam := context.Query("page")

	limit := 3
	page := 1

	if param != "" {
		l, err := strconv.Atoi(param)
		if err == nil && l > 0 {
			limit = l
		}
	}

	if pageParam != "" {
		p, err := strconv.Atoi(pageParam)
		if err == nil && p > 0 {
			page = p
		}
	}

	offset := (page - 1) * limit

	photos, err := Model.FindPhotosByContextPagination("gallery", limit, offset)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get photos", "error": err.Error()})
		return
	}

	count, err := Model.CountPhotosByContext("gallery")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get photo count", "error": err.Error()})
		return
	}

	totalPages := int(math.Ceil(float64(count) / float64(limit)))

	response := gin.H{
		"message":     "Gallery retrieved successfully!",
		"photos":      photos,
		"currentPage": page,
		"totalPages":  totalPages,
	}

	if page > 1 {
		response["previousPage"] = page - 1
	}

	if page < totalPages {
		response["nextPage"] = page + 1
	}

	context.JSON(http.StatusOK, response)
}

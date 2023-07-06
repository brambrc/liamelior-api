package Controller

import (
	"liamelior-api/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Setlist          string `json:"setlist" binding:"required"`
	showDate         string `json:"show_date" binding:"required"`
	time             string `json:"time" binding:"required"`
	totalPerformance int    `json:"total_performance" binding:"required"`
}

func GetUpcomingAndPastShows(c *gin.Context) {
	// Retrieve upcoming shows from the database
	upcomingShows, err := Model.GetUpcomingShows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve upcoming shows",
		})
		return
	}

	pastShows, err := Model.GetPastShowCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve past shows",
		})
		return
	}

	if len(upcomingShows) == 0 && len(pastShows) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":        "success",
			"upcoming_shows": "No upcoming shows",
			"past_shows":     "No past shows",
		})
		return
	} else if len(upcomingShows) == 0 && len(pastShows) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":        "success",
			"upcoming_shows": "No upcoming shows",
			"past_shows":     pastShows,
		})
		return
	} else if len(upcomingShows) != 0 && len(pastShows) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":        "success",
			"upcoming_shows": upcomingShows,
			"past_shows":     "No past shows",
		})
		return
	} else {
		// Create the response
		response := gin.H{
			"message":        "success",
			"upcoming_shows": upcomingShows,
			"past_shows":     pastShows,
		}

		c.JSON(http.StatusOK, response)
		return

	}

}

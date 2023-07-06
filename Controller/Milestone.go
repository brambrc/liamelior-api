package Controller

import (
	"liamelior-api/Database"
	"liamelior-api/Model"

	"github.com/gin-gonic/gin"
)

type Milestone struct {
	Milestone string `json:"milestone" binding:"required"`
	Place     string `json:"place" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Svg       string `json:"svg" binding:"required"`
}

type MilestoneResponse struct {
	ID        uint   `json:"ID"`
	Milestone string `json:"milestone" binding:"required"`
	Place     string `json:"place" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Svg       string `json:"svg" binding:"required"`
}

func SaveMilestone(context *gin.Context) {
	var input Model.Milestone
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(400, gin.H{"message": "Failed to bind JSON", "error": err.Error()})
		return
	}
	milestone := Model.Milestone{
		Milestone: input.Milestone,
		Place:     input.Place,
		Date:      input.Date,
		Svg:       input.Svg,
	}

	_, err = milestone.Save()

	if err != nil {
		context.JSON(400, gin.H{"message": "Failed to create milestone", "error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Milestone created successfully"})

}

func GetMilestone(context *gin.Context) {
	var milestones []Model.Milestone
	var milestoneResponses []MilestoneResponse

	err := Database.Database.Find(&milestones).Error

	if err != nil {
		context.JSON(400, gin.H{"message": "Failed to fetch milestone", "error": err.Error()})
		return
	}

	for _, milestone := range milestones {
		milestoneResponses = append(milestoneResponses, MilestoneResponse{
			ID:        milestone.ID,
			Milestone: milestone.Milestone,
			Place:     milestone.Place,
			Date:      milestone.Date,
			Svg:       milestone.Svg,
		})
	}

	context.JSON(200, gin.H{"message": "Milestone fetched successfully", "data": milestoneResponses})
}

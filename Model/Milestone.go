package Model

import (
	"liamelior-api/Database"

	"gorm.io/gorm"
)

type Milestone struct {
	gorm.Model
	Milestone string `json:"milestone" binding:"required"`
	Place     string `json:"place" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Svg       string `json:"svg" binding:"required"`
}

type MilestoneResponse struct {
	gorm.Model
	ID        uint   `json:"ID"`
	Milestone string `json:"milestone" binding:"required"`
	Place     string `json:"place" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Svg       string `json:"svg" binding:"required"`
}

func (m *Milestone) Save() (*Milestone, error) {

	var err error
	err = Database.Database.Create(&m).Error

	if err != nil {
		return &Milestone{}, err
	}

	return m, nil
}

func (m *Milestone) Find() (*[]MilestoneResponse, error) {
	var err error
	var milestones []Milestone
	var milestoneResponses []MilestoneResponse

	err = Database.Database.Find(&milestones).Error

	if err != nil {
		return &[]MilestoneResponse{}, err
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

	return &milestoneResponses, nil
}

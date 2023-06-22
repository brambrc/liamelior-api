package Model

import (
	"liamelior-api/Database"

	"gorm.io/gorm"
)

type TextContent struct {
	gorm.Model
	Title         string `json:"title"`
	SubHeader     string `json:"sub_header"`
	CaraouselText string `json:"caraousel_text"`
	About         string `json:"about"`
	LinkAbout     string `json:"link_about"`
	MilestoneText string `json:"milestone_text"`
	ShowsText     string `json:"shows_text"`
	FooterText    string `json:"footer_text"`
	InstagramLink string `json:"instagram_link"`
	TwitterLink   string `json:"twitter_link"`
	TiktokLink    string `json:"tiktok_link"`
	CopyRight     string `json:"copy_right"`
	TabIcon       string `json:"tab_icon"`
	TabText       string `json:"tab_text"`
	HeroImage     string `json:"hero_image"`
	FooterImage   string `json:"footer_image"`
}

func (textContent *TextContent) Save() (*TextContent, error) {
	var err error

	err = Database.Database.Create(&textContent).Error
	if err != nil {
		return &TextContent{}, err
	}

	return textContent, nil
}

func FindTextContentById() (*TextContent, error) {
	var err error
	var textContent TextContent

	err = Database.Database.First(&textContent, 1).Error
	if err != nil {
		return &TextContent{}, err
	}

	return &textContent, nil
}

func (textContent *TextContent) Update() (*TextContent, error) {
	var err error

	err = Database.Database.Save(&textContent).Error
	if err != nil {
		return &TextContent{}, err
	}

	return textContent, nil
}

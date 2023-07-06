package Controller

import (
	"liamelior-api/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TextContent struct {
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

func TextContentStore(context *gin.Context) {
	var input Model.TextContent
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to bind JSON", "error": err.Error()})
		return
	}
	textContent := Model.TextContent{
		Title:         input.Title,
		SubHeader:     input.SubHeader,
		CaraouselText: input.CaraouselText,
		About:         input.About,
		LinkAbout:     input.LinkAbout,
		MilestoneText: input.MilestoneText,
		ShowsText:     input.ShowsText,
		FooterText:    input.FooterText,
		InstagramLink: input.InstagramLink,
		TwitterLink:   input.TwitterLink,
		TiktokLink:    input.TiktokLink,
		CopyRight:     input.CopyRight,
		TabIcon:       input.TabIcon,
		TabText:       input.TabText,
		HeroImage:     input.HeroImage,
		FooterImage:   input.FooterImage,
	}
	_, err = textContent.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create text content", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Text content created successfully!", "textContent": textContent})
}

func GetTextContent(context *gin.Context) {
	textContent, err := Model.FindTextContentById()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to find text content", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Text content found", "textContent": textContent})
}

func TextContentUpdate(context *gin.Context) {
	var input Model.TextContent
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to bind JSON", "error": err.Error()})
		return
	}

	//findData := Model.TextContent{}

	findData, err := Model.FindTextContentById()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to find text content", "error": err.Error()})
		return
	}

	//update data

	findData.Title = input.Title
	findData.SubHeader = input.SubHeader
	findData.CaraouselText = input.CaraouselText
	findData.About = input.About
	findData.LinkAbout = input.LinkAbout
	findData.MilestoneText = input.MilestoneText
	findData.ShowsText = input.ShowsText
	findData.FooterText = input.FooterText
	findData.InstagramLink = input.InstagramLink
	findData.TwitterLink = input.TwitterLink
	findData.TiktokLink = input.TiktokLink
	findData.CopyRight = input.CopyRight
	findData.TabIcon = input.TabIcon
	findData.TabText = input.TabText
	findData.HeroImage = input.HeroImage
	findData.FooterImage = input.FooterImage

	textContent, err := findData.Update()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update text content", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Text content updated successfully!", "textContent": textContent})
}

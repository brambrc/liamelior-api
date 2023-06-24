package Controller


import (
	"liamelior-api/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Milestone struct {
	Milestone string `json: "milestone" binding: "required"`
	Place string `json: "place" binding: "required"`
	Date string `json"`
}



func SaveMilestone(context *gin.Context) {

}
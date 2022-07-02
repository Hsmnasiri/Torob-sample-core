package api

import (
	"net/http"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/gin-gonic/gin"
)

type CreateReportInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func CreateReport(c *gin.Context) {
	var input CreateReportInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := entity.Report{}

	r.Name = input.Name
	r.Description = input.Description
	//p.Types = input.Types

	rout, err := r.SaveReport()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "product": rout})

}
func GetReports(c *gin.Context) {

}
func GetOneReport(c *gin.Context) {

}

func UpdateReport(c *gin.Context) {

}

func DeleteReport(c *gin.Context) {

}

package api

import (
	"net/http"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/gin-gonic/gin"
)

type CreateTypeInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateTypes(c *gin.Context) {
	var input CreateTypeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := entity.Product{}

	t.Name = input.Name
	//p.Types = input.Types

	tout, err := t.SaveProduct()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "product": tout})
}

func GetTypes(c *gin.Context) {

	types, err := entity.GetTypes()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "users find success", "types": types})

}

package api

import (
	"fmt"
	"net/http"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/Hsmnasiri/Torob-sample-core/utils"
	"github.com/gin-gonic/gin"
)

type CreateSubTypeInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateSubTypes(c *gin.Context) {
	var input CreateSubTypeInput
	user_role, err := utils.ExtractTokenRole(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user_role)
	if user_role != "admin" {
		fmt.Println("we are in error handler")
		c.JSON(http.StatusForbidden, gin.H{"error": "unauthorized"})
		return
	}
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := entity.SubType{}

	t.Name = input.Name

	tout, err := t.SaveSubType()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create Type success", "type": tout})
}

func GetSubTypes(c *gin.Context) {

	subTypes, err := entity.GetSubTypes()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "create Type success", "subTypes": subTypes})

}

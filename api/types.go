package api

import (
	"fmt"
	"net/http"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/Hsmnasiri/Torob-sample-core/utils"
	"github.com/gin-gonic/gin"
)

type CreateTypeInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateTypes(c *gin.Context) {
	var input CreateTypeInput
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

	t := entity.Type{}

	t.Name = input.Name
	//p.Types = input.Types

	tout, err := t.SaveType()

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

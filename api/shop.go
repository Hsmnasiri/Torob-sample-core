package api

import (
	"net/http"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/gin-gonic/gin"
)

type CreateShopInput struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

func CreateShop(c *gin.Context) {
	var input CreateShopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := entity.Shop{}

	s.Name = input.Name
	s.PhoneNumber = input.Phone
	//p.Types = input.Types

	Sout, err := s.SaveShop()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "product": Sout})

}

func UpdateShop(c *gin.Context) {
	var input CreateShopInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s := entity.Shop{}

	s.Name = input.Name
	s.PhoneNumber = input.Phone
	//p.Types = input.Types

	Sout, err := s.SaveShop()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "product": Sout})

}

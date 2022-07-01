package api

import (
	"net/http"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Name  string `json:"name" binding:"required"`
	Price string `json:"price" binding:"required"`
	Types string `json:"types"  `
}

func CreateProduct(c *gin.Context) {
	var input CreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := entity.Product{}

	p.Name = input.Name
	p.Price = input.Price
	//p.Types = input.Types

	pout, err := p.SaveProduct()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "product": pout})

}


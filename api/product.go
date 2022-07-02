package api

import (
	"fmt"
	"net/http"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Name         string `json:"name" binding:"required"`
	LowestPrice  string `json:"lowest_price" binding:"required"`
	HighestPrice string `json:"highest_price" binding:"required"`
	Types        string `json:"types"`
	Img          string `json:"img"`
}

func CreateProduct(c *gin.Context) {
	var input CreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p := entity.Product{}

	p.Name = input.Name
	p.HighestPrice = input.HighestPrice
	p.LowestPrice = input.LowestPrice

	pout, err := p.SaveProduct()
	r := entity.DB.Model(&entity.Type{}).Where("ID = ?", input.Types).Association("Products").Append(&pout)

	fmt.Println(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "product": pout})

}

func GetProducts(c *gin.Context) {

}
func GetOneProduct(c *gin.Context) {

}

func UpdateProducts(c *gin.Context) {

}

func DeleteProducts(c *gin.Context) {

}

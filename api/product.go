package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Name         string `json:"name" binding:"required"`
	LowestPrice  string `json:"lowest_price" binding:"required"`
	HighestPrice string `json:"highest_price" binding:"required"`
	TypeID       uint   `json:"type_id"`
	SubTypeID    uint   `json:"subtype_id"`
	Img          string `json:"img"`
	Shop         uint   `json:"shop"`
}
type IncShop struct {
	Shop uint `json:"shop" binding:"required"`
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
	p.Img = input.Img
	p.TypeID = input.TypeID
	p.SubTypeID = input.SubTypeID

	pout, err := p.SaveProduct()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "product": pout})

}

func GetProducts(c *gin.Context) {
	products, err := entity.GetProducts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "products find success", "products": products})

}
func GetOneProduct(c *gin.Context) {

}

func UpdateProducts(c *gin.Context) {

}

func DeleteProducts(c *gin.Context) {

}
func NewShopForProduct(c *gin.Context) {
	var input IncShop

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pid, _ := strconv.ParseUint(fmt.Sprintf("%.0f", c.Param("productId")), 10, 32)
	err := entity.IncrementShop(input.Shop, uint(pid))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "shop incremented success"})
}

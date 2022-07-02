package main

import (
	"github.com/Hsmnasiri/Torob-sample-core/api"
	"github.com/Hsmnasiri/Torob-sample-core/entity"
	"github.com/Hsmnasiri/Torob-sample-core/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.ConnectDataBase()

	r := gin.Default()
	public := r.Group("/v1")

	public.POST("/register", api.Register)
	public.POST("/login", api.Login)

	protected := r.Group("/v1/api")
	protected.Use(utils.JwtAuthMiddleware())
	protected.GET("/user", api.CurrentUser)
	protected.GET("/users", api.GetUsers)

	shopApi := protected.Group("/shop")
	protected.POST("/shop", api.CreateShop)
	shopApi.PUT("/", api.UpdateShop)
	shopApi.GET("/", api.CreateProduct)
	shopApi.GET("/:shopID", api.CreateProduct)
	shopApi.DELETE("/", api.UpdateShop)

	productApi := protected.Group("/product")
	productApi.POST("/", api.CreateProduct)
	productApi.GET("/", api.GetProducts)
	productApi.GET("/:productId", api.GetOneProduct)
	productApi.PUT("/:productId", api.UpdateProducts)
	productApi.PUT("/shop/:productId", api.NewShopForProduct)
	productApi.DELETE("/", api.DeleteProducts)

	typeApi := protected.Group("/types")
	typeApi.GET("/", api.GetTypes)
	typeApi.POST("/", api.CreateTypes)

	subtypeApi := protected.Group("/subtypes")
	subtypeApi.GET("/", api.GetSubTypes)
	subtypeApi.POST("/", api.CreateSubTypes)

	reportApi := protected.Group("/report")
	reportApi.POST("/", api.CreateReport)
	reportApi.GET("/:shopId", api.GetReports)

	r.Run(":9090")

}

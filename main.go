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
	public := r.Group("/api")

	public.POST("/register", api.Register)
	public.POST("/login", api.Login)
	public.POST("/product", api.CreateProduct)
	public.POST("/types",utils.JwtAuthMiddleware(), api.CreateTypes)
	public.POST("/shop", api.CreateShop)

	public.PUT("/shop", api.UpdateShop)

	public.GET("/shop", api.CreateProduct)
	public.GET("/user", api.GetUsers)
	public.GET("/product", api.GetTypes)
	public.GET("/types", api.GetTypes)

	protected := r.Group("/api/admin")
	protected.Use(utils.JwtAuthMiddleware())
	protected.GET("/user", api.CurrentUser)

	r.Run(":9090")

}

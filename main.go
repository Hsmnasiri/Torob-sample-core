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

	protected := r.Group("/api/admin")
	protected.Use(utils.JwtAuthMiddleware())
	protected.GET("/user", api.CurrentUser)

	r.Run(":9090")

}

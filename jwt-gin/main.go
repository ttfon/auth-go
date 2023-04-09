package main

import (
	"jwt-go/controller"
	"jwt-go/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controller.Register)
	public.POST("/login", controller.Login)
	r.Run(":8080")
}

package main

import (
	"go-api/controller"
	"go-api/usecase"
	"github.com/gin-gonic/gin"
)

func main(){

	server := gin.Default()

	ProductUsecase := usecase.NewProductUseCase()

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	 
	server.GET("/product", ProductController.GetProducts)
	server.Run(":3000")

}
package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)

	}
	
    //repository
	ProductRepository := repository.NewProductRepository(dbConnection)
    //usecase
	ProductUsecase := usecase.NewProductUseCase(ProductRepository)
    //controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/product", ProductController.GetProducts)
	server.Run(":3000")

}

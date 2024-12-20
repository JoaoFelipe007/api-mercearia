package main

import (
	"api-mercearia-go/controller"
	"api-mercearia-go/db"
	"api-mercearia-go/repository"
	"api-mercearia-go/usecase"

	"github.com/gin-gonic/gin"
)

// @title API MERCEARIA
// @version 1.0
// @description This is a RESTful API in Go using Swagger
// @contact.name João FElipe
// @joaofelipe.developer@gmail.com
// @host localhost:8080
// @BasePath /v1
func main() {
	router := gin.Default()

	// iniciando conexão do banco de dados
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	//pacotes da repository
	CategoryRepository := repository.NewCategoryRepository(dbConnection)

	//pacote usecase
	CategoryUsecase := usecase.NewCategoryUsecase(CategoryRepository)

	//pacote controlller
	CategoryControler := controller.NewCategoryControler(CategoryUsecase)

	router.GET("/categories", CategoryControler.GetCategories)
	router.GET("/category/:id", CategoryControler.GetCategoryById)
	router.POST("/category", CategoryControler.CreateCategory)
	router.DELETE("/category/:id", CategoryControler.DeleteCategory)

	router.Run("localhost:8080")
}

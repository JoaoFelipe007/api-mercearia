package main

import (
	"api-mercearia-go/controller"
	"api-mercearia-go/db"
	"api-mercearia-go/repository"
	"api-mercearia-go/usecase"

	"github.com/gin-gonic/gin"
)

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
	router.POST("/category", CategoryControler.CreateCategory)
	router.DELETE("/category/:id", CategoryControler.DeleteCategoty)

	router.Run("localhost:8080")
}

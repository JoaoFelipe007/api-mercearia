package main

import (
	"api-mercearia-go/controller"
	"api-mercearia-go/db"
	_ "api-mercearia-go/docs"
	"api-mercearia-go/repository"
	"api-mercearia-go/usecase"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API MERCEARIA
// @version 1.0
// @description This is a RESTful API in Go using Swagger
// @contact.name João FElipe
// @joaofelipe.developer@gmail.com
// @host localhost:8080
// @BasePath /
func main() {
	// Iniciando o Gin Router
	router := gin.Default()

	// Configurando o Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciando a conexão com o banco de dados
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err) // Caso haja erro, o servidor falha e retorna o erro
	}

	// Criando instâncias dos pacotes da camada de repositório, usecase e controller
	CategoryRepository := repository.NewCategoryRepository(dbConnection)
	CategoryUsecase := usecase.NewCategoryUsecase(CategoryRepository)
	CategoryControler := controller.NewCategoryControler(CategoryUsecase)

	// Definindo as rotas
	router.GET("/categories", CategoryControler.GetCategories)
	router.GET("/category/:id", CategoryControler.GetCategoryById)
	router.POST("/category", CategoryControler.CreateCategory)
	router.DELETE("/category/:id", CategoryControler.DeleteCategory)

	// Rodando o servidor na porta 8080
	router.Run(":8080")
}

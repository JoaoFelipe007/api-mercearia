package main

import (
	"api-mercearia-go/controller"
	"api-mercearia-go/db"
	_ "api-mercearia-go/docs"
	"api-mercearia-go/middleware"
	"api-mercearia-go/repository"
	"api-mercearia-go/usecase"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GROCERY API
// @version 1.0
// @description This is a RESTful API in Go using Swagger
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @contact.name João Felipe
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
	defer dbConnection.Close() // Feche o banco ao final da execução da aplicação

	// Criando instâncias dos pacotes da camada de repositório, usecase e controller
	CategoryRepository := repository.NewCategoryRepository(dbConnection)
	CategoryUsecase := usecase.NewCategoryUsecase(CategoryRepository)
	CategoryControler := controller.NewCategoryControler(CategoryUsecase)

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	PersonRepository := repository.NewPersonRepository(dbConnection)
	PersonUseCase := usecase.NewPersonUsecase(PersonRepository)
	PersonController := controller.NewPersonController(PersonUseCase)

	AutenticationUseCase := usecase.NewAutenticationUsecase(PersonRepository)
	AutenticationController := controller.NewAutenticationControler(AutenticationUseCase)

	// Definindo as rotas
	categoryRoutes := router.Group("/category")
	categoryRoutes.Use(middleware.AuthenticationMiddleware())
	{
		categoryRoutes.GET("/:id", CategoryControler.GetCategoryById)
		categoryRoutes.POST("/category", CategoryControler.CreateCategory)
		categoryRoutes.DELETE("/:id", CategoryControler.DeleteCategory)
		categoryRoutes.PUT("/change-status/:id", CategoryControler.ChangeStatus)
	}

	categoriesRoutes := router.Group("/categories")
	categoriesRoutes.Use(middleware.AuthenticationMiddleware())
	{
		categoryRoutes.GET("/", CategoryControler.GetCategories)
	}

	productRoutes := router.Group("/product")
	productRoutes.Use(middleware.AuthenticationMiddleware())
	{
		productRoutes.POST("/", ProductController.CreateProduct)
		productRoutes.GET("/:id", ProductController.GetProductById)
		productRoutes.PUT("/", ProductController.ChangeProduct)
		productRoutes.DELETE("/:id", ProductController.DeleteProduct)

	}

	productsRoutes := router.Group("/products")
	productsRoutes.Use(middleware.AuthenticationMiddleware())
	{
		productsRoutes.GET("/", ProductController.GetProducts)
	}
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", PersonController.CreatePerson)
		authRoutes.POST("/authorization", AutenticationController.Authorization)
	}
	// Rodando o servidor na porta 8080
	router.Run(":8080")
}

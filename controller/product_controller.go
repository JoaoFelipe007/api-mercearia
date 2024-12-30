package controller

import (
	"api-mercearia-go/model"
	"api-mercearia-go/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(productUsecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: productUsecase,
	}
}

// GetProducts godoc
// @Summary list all products
// @Description Returns all active products or not
// @Tags Product
// @Produce json
// @Success 200 {array} model.Product
// @Router /products [get]
func (pc *ProductController) GetProducts(ctx *gin.Context) {
	productList, err := pc.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, productList)
}

// GetProducts godoc
// @Summary list product by id
// @Description Returns the specified product
// @Tags Product
// @Produce json
// @Param id path int true "ID of Product"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Response "Erro de validação nos dados"
// @Router /product/{id} [get]
func (pc *ProductController) GetProductById(ctx *gin.Context) {
	idString := ctx.Param("id")

	if idString == "" {
		response := model.Response{Message: "The id needs to be provided"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(idString)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	product, err := pc.productUsecase.GetProductById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, product)

}

// CreateProduct godoc
// @Summary Save a product
// @Description Save a product
// @Tags Product
// @Produce json
// @Param category body model.Product true "Dados do novo produto"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Response "Erro de validação nos dados"
// @Router /product [post]
func (pc *ProductController) CreateProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(product)

	if product.Description == "" {
		reponse := model.Response{Message: "The description needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	if product.Quantity == 0.0 {
		reponse := model.Response{Message: "The quantity needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	if float64(product.CategoryID) == 0 {
		reponse := model.Response{Message: "The categoria needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	productResult, err := pc.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, productResult)
}

// CreateProduct godoc
// @Summary Save a product
// @Description Save a product
// @Tags Product
// @Produce json
// @Param category body model.Product true "Dados do produto existente"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.Response "Erro de validação nos dados"
// @Router /product [put]
func (pc *ProductController) ChangeProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(product)

	if product.ID == 0 {
		reponse := model.Response{Message: "The ID needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	if product.Description == "" {
		reponse := model.Response{Message: "The description needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	if product.Quantity == 0.0 {
		reponse := model.Response{Message: "The quantity needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	if float64(product.CategoryID) == 0 {
		reponse := model.Response{Message: "The categoria needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	productResult, err := pc.productUsecase.ChangeProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, productResult)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Deletes the specified product
// @Tags Product
// @Produce json
// @Param id path int true "ID of product"
// @Success 200 {string} string "Delatado com sucesso"
// @Router /product/{id} [delete]
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	var idProduct int

	idStr := ctx.Param("id") // Obtem o valor do parâmetro como string
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	idProduct = id

	messageSucces, err := c.productUsecase.DeleteProduct(idProduct)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": messageSucces,
	})
}

package controller

import (
	"api-mercearia-go/model"
	"api-mercearia-go/usecase"
	"fmt"
	"net/http"

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

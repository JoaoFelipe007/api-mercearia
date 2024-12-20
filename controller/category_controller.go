package controller

import (
	"api-mercearia-go/model"
	"api-mercearia-go/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryControler(usecase usecase.CategoryUsecase) categoryController {
	return categoryController{
		categoryUsecase: usecase,
	}
}

// GetCategories godoc
// @Summary Lista todas as categorias
// @Description Retorna todas as categorias disponíveis no sistema.
// @Tags Categorias
// @Produce json
// @Success 200 {array} model.Category
// @Router /categories [get]
func (c *categoryController) GetCategories(ctx *gin.Context) {

	categories, err := c.categoryUsecase.GetCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, categories)

}

func (c *categoryController) GetCategoryById(ctx *gin.Context) {

	idStr := ctx.Param("id") // Obtem o valor do parâmetro como string

	if idStr == "" {
		response := model.Response{Message: "O id não pode ser nulo"}
		ctx.JSON(http.StatusBadRequest, response)
	}

	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	category, err := c.categoryUsecase.GetCategoryById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, category)

}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var category model.Category
	err := ctx.BindJSON(&category)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertCategory, err := c.categoryUsecase.CreateCategory(category)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertCategory)
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	var idCategory int

	idStr := ctx.Param("id") // Obtem o valor do parâmetro como string
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	idCategory = id

	messageSucces, err := c.categoryUsecase.DeleteCategory(idCategory)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": messageSucces,
	})
}

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
// @Summary List all categories
// @Description Returns all categories available in the system.
// @Tags Category
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

// GetCategoryById godoc
// @Summary List the category by id
// @Description Returns the specified category
// @Tags Category
// @Produce json
// @Param id path int true "ID of Category"
// @Success 200 {object} model.Category
// @Router /category/{id} [get]
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

// CreateCategory
// @Summary Create a new category
// @Description Create a new category
// @Tags Category
// @Produce json
// @Param category body model.Category true "Dados da nova categoria"
// @Success 200 {object} model.Category
// @Router /category [post]
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

// DeleteCategory godoc
// @Summary Delete a category
// @Description Deletes the specified category
// @Tags Category
// @Produce json
// @Param id path int true "ID da Categoria"
// @Success 200 {string} string "Delatado com sucesso"
// @Router /category/{id} [delete]
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

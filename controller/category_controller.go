package controller

import (
	"api-mercearia-go/usecase"
	"net/http"

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

func (c *categoryController) GetCategories(ctx *gin.Context) {

	categories, err := c.categoryUsecase.GetCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, categories)

}

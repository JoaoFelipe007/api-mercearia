package controller

import (
	"api-mercearia-go/model"
	"api-mercearia-go/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AutenticationController struct {
	autenticationUsecase usecase.AutenticationUsecase
}

func NewAutenticationControler(usecase usecase.AutenticationUsecase) AutenticationController {
	return AutenticationController{
		autenticationUsecase: usecase,
	}
}

// Authorization godoc
// @Summary Generate token JWT
// @Description Generate token JWT
// @Tags Authorization
// @Produce json
// @Param category body model.UserLogin true "Dados para login"
// @Success 200 {object} model.UserLogin
// @Failure 400 {object} model.Response "Erro de validação nos dados"
// @Failure 406 {object} model.Response "Erro de validação nos dados"
// @Router /authorization [post]
func (ac *AutenticationController) Authorization(ctx *gin.Context) {
	var userLogin model.UserLogin

	ctx.BindJSON(&userLogin)

	if userLogin.Username == "" {
		ctx.JSON(http.StatusNotAcceptable, model.Response{Message: "Username needs to be entered"})
		return
	}
	if userLogin.Password == "" {
		ctx.JSON(http.StatusNotAcceptable, model.Response{Message: "Password needs to be entered"})
		return
	}

	token, err := ac.autenticationUsecase.Login(userLogin)

	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if token == "" {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "Login or password is incorrect"})
		return
	}

	ctx.Header("Authorization", token)
	ctx.JSON(http.StatusOK, token)
}

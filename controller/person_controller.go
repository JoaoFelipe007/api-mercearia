package controller

import (
	"api-mercearia-go/model"
	"api-mercearia-go/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonController struct {
	personUsecase usecase.PersonUsecase
}

func NewPersonController(personUsecase usecase.PersonUsecase) PersonController {
	return PersonController{
		personUsecase: personUsecase,
	}
}

// CreatePerson godoc
// @Summary Save a person
// @Description Save a person
// @Tags Authorization
// @Produce json
// @Param category body model.Person true "Dados do novo usuario"
// @Success 200 {object} model.Person
// @Failure 400 {object} model.Response "Erro de validação nos dados"
// @Router /auth/register [post]
func (pc *PersonController) CreatePerson(ctx *gin.Context) {
	var person model.Person

	err := ctx.BindJSON(&person)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if person.Name == "" {
		reponse := model.Response{Message: "The name needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	if person.Password == "" {
		reponse := model.Response{Message: "The password needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	if person.Email == "" {
		reponse := model.Response{Message: "The email needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	if person.Document == "" {
		reponse := model.Response{Message: "The document needs to be informed!"}
		ctx.JSON(http.StatusBadRequest, reponse)
		return
	}

	personResult, err := pc.personUsecase.CreatePerson(person)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, personResult)

}

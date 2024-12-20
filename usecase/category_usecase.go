package usecase

import (
	"api-mercearia-go/model"
	"api-mercearia-go/repository"
	"fmt"
)

type CategoryUsecase struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepository repository.CategoryRepository) CategoryUsecase {
	return CategoryUsecase{
		categoryRepository: categoryRepository,
	}
}

func (cu *CategoryUsecase) GetCategories() ([]model.Category, error) {
	return cu.categoryRepository.GetCategories()
}

func (cu *CategoryUsecase) GetCategoryById(id int) (model.Category, error) {

	category, err := cu.categoryRepository.GetCategoryById(id)

	if err != nil {
		fmt.Print(err)
		return model.Category{}, err
	}

	return category, nil
}

func (cu *CategoryUsecase) CreateCategory(category model.Category) (model.Category, error) {

	categoryId, err := cu.categoryRepository.CreateCategory(category)

	if err != nil {
		fmt.Print(err)
		return model.Category{}, err
	}
	category.ID = categoryId
	return category, nil
}

func (cu *CategoryUsecase) ChangeStatus(id int) (model.Category, error) {
	category, err := cu.categoryRepository.ChangeStatus(id)

	if err != nil {
		fmt.Print(err)
		return model.Category{}, err
	}

	return category, nil
}

func (cu *CategoryUsecase) DeleteCategory(id int) (string, error) {

	messageSucces, err := cu.categoryRepository.DeleteCategoty(id)

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return messageSucces, nil

}

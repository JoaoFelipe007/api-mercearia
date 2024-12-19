package usecase

import (
	"api-mercearia-go/model"
	"api-mercearia-go/repository"
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

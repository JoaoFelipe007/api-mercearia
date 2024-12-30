package usecase

import (
	"api-mercearia-go/model"
	"api-mercearia-go/repository"
	"fmt"
)

type ProductUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		productRepository: productRepository,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.productRepository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productResult, err := pu.productRepository.CreateProduct(product)

	if err != nil {
		fmt.Print(err)
		return model.Product{}, err
	}

	return productResult, nil
}

package usecase

import (
	"api-mercearia-go/model"
	"api-mercearia-go/repository"
	"fmt"
	"log"
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

func (pu *ProductUsecase) ChangeProduct(product model.Product) (model.Product, error) {

	productResult, err := pu.productRepository.ChangeProduct(product)

	if err != nil {
		fmt.Print(err)
		return model.Product{}, err
	}

	return productResult, nil
}

func (pu *ProductUsecase) GetProductById(id int) (model.Product, error) {
	productResult, err := pu.productRepository.GetProductById(id)

	if err != nil {
		log.Fatalln(err)
		return model.Product{}, err
	}

	return productResult, nil

}

func (cu *ProductUsecase) DeleteProduct(id int) (string, error) {

	messageSucces, err := cu.productRepository.DeleteProduct(id)

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return messageSucces, nil

}

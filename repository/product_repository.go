package repository

import (
	"api-mercearia-go/model"
	"database/sql"
	"fmt"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM product"

	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Print(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Description,
			&productObj.Priority,
			&productObj.Price,
			&productObj.Quantity,
			&productObj.CategoryID,
			&productObj.Status,
			&productObj.RegistrationDate,
		)

		if err != nil {
			fmt.Print(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()
	return productList, err
}

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	var productResult model.Product

	query, err := pr.connection.Prepare("INSERT INTO product (description, priority, price, quantity, id_category) " +
		" VALUES ($1,$2,$3,$4,$5) returning *")
	if err != nil {
		fmt.Print(err)
		return model.Product{}, err
	}
	fmt.Print(product)
	err = query.QueryRow(product.Description, product.Priority, product.Price, product.Quantity, product.CategoryID).Scan(
		&productResult.ID,
		&productResult.CategoryID,
		&productResult.RegistrationDate,
		&productResult.Price,
		&productResult.Description,
		&productResult.Quantity,
		&productResult.Status,
		&productResult.Priority,
	)
	if err != nil {
		fmt.Print(err)
		return model.Product{}, err
	}

	query.Close()
	return productResult, err
}

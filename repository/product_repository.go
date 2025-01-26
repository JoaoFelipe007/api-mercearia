package repository

import (
	"api-mercearia-go/model"
	"database/sql"
	"fmt"
	"log"
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
			&productObj.DateChange,
		)

		if err != nil {
			fmt.Print(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

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
		&productResult.Description,
		&productResult.Priority,
		&productResult.Price,
		&productResult.Quantity,
		&productResult.CategoryID,
		&productResult.Status,
		&productResult.RegistrationDate,
		&productResult.DateChange,
	)
	if err != nil {
		fmt.Print(err)
		return model.Product{}, err
	}

	return productResult, err
}

func (pr *ProductRepository) GetProductById(id int) (model.Product, error) {
	var product model.Product

	query := "SELECT * FROM product WHERE id = $1"

	err := pr.connection.QueryRow(query, id).Scan(
		&product.ID,
		&product.Description,
		&product.Priority,
		&product.Price,
		&product.Quantity,
		&product.CategoryID,
		&product.Status,
		&product.RegistrationDate,
		&product.DateChange,
	)
	if err != nil {

		if err == sql.ErrNoRows { //nenhuma liha encontrada
			return model.Product{}, nil
		}
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) ChangeProduct(product model.Product) (model.Product, error) {
	var productResult model.Product

	query, err := pr.connection.Prepare("Update product set " +
		"description = $1," +
		"priority = $2, " +
		"price = $3," +
		"quantity = $4," +
		"id_category = $5," +
		"status = $6," +
		"date_change = now()" +
		"WHERE id = $7 " +
		"returning *")
	if err != nil {
		log.Fatal(err)
		return model.Product{}, err
	}

	err = query.QueryRow(product.Description, product.Priority, product.Price, product.Quantity, product.CategoryID, product.Status, product.ID).Scan(
		&productResult.ID,
		&productResult.Description,
		&productResult.Priority,
		&productResult.Price,
		&productResult.Quantity,
		&productResult.CategoryID,
		&productResult.Status,
		&productResult.RegistrationDate,
		&productResult.DateChange,
	)

	if err != nil {
		if err == sql.ErrNoRows { //nenhuma liha encontrada
			return model.Product{}, nil
		}
		return model.Product{}, err
	}

	return productResult, nil

}

func (cr *ProductRepository) DeleteProduct(id int) (string, error) {

	query, err := cr.connection.Prepare("delete from product where id = $1")

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	result, err := query.Exec(id)

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		fmt.Print(err)
		return "", err
	}

	if rowsAffected == 0 {
		return "", fmt.Errorf("nenhuma linha foi deletada")
	}

	return "Delatado com sucesso", nil
}

package repository

import (
	"api-mercearia-go/model"
	"database/sql"
	"fmt"
)

type CategoryRepository struct {
	connection *sql.DB
}

func NewCategoryRepository(connection *sql.DB) CategoryRepository {
	return CategoryRepository{
		connection: connection,
	}
}

func (pr *CategoryRepository) GetCategories() ([]model.Category, error) { // O que seria o(pr *)
	query := "select id, description, priority from category "

	rows, err := pr.connection.Query(query)

	if err != nil {
		fmt.Print(err)
		return []model.Category{}, err
	}

	var categoryList []model.Category
	var categoryObj model.Category

	for rows.Next() {
		err = rows.Scan(
			&categoryObj.ID, //Endereço de memoria
			&categoryObj.Description,
			&categoryObj.Priority,
		)

		if err != nil {
			fmt.Print(err)
			return []model.Category{}, err
		}

		categoryList = append(categoryList, categoryObj)
	}
	rows.Close()

	return categoryList, err
}

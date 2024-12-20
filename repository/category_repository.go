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

func (cr *CategoryRepository) GetCategoryById(id int) (model.Category, error) {
	query := "SELECT * FROM category WHERE id = $1"
	db := cr.connection
	defer db.Close() // Garante que a conexão será fechada ao final da função

	var category model.Category

	err := db.QueryRow(query, id).Scan(&category.ID, &category.Description, &category.Priority)

	if err != nil {
		if err == sql.ErrNoRows { //nenhuma liha encontrada
			return model.Category{}, nil
		}
		return model.Category{}, err
	}
	return category, err
}

func (cr *CategoryRepository) CreateCategory(category model.Category) (int, error) {

	var id int

	query, err := cr.connection.Prepare("INSERT INTO category " +
		"(description, priority) " +
		"VALUES ($1,$2) returning id")

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	err = query.QueryRow(category.Description, category.Priority).Scan(&id)

	if err != nil {
		fmt.Print(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (cr *CategoryRepository) DeleteCategoty(id int) (string, error) {

	query, err := cr.connection.Prepare("delete from category where id = $1")

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
	query.Close()

	return "Delatado com sucesso", nil
}

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

	return categoryList, err
}

func (cr *CategoryRepository) GetCategoryById(id int) (model.Category, error) {
	query := "SELECT * FROM category WHERE id = $1"

	db := cr.connection

	var category model.Category

	err := db.QueryRow(query, id).Scan(&category.ID, &category.Description, &category.Priority, &category.Status)

	if err != nil {
		if err == sql.ErrNoRows { //nenhuma liha encontrada
			return model.Category{}, nil
		}
		return model.Category{}, err
	}
	return category, err
}

func (cr *CategoryRepository) CreateCategory(category model.Category) (model.Category, error) {

	var categoryResult model.Category

	query, err := cr.connection.Prepare("INSERT INTO category (description, priority) " +
		"VALUES ($1,$2) returning *")

	if err != nil {
		fmt.Print(err)
		return model.Category{}, err
	}

	err = query.QueryRow(category.Description, category.Priority).Scan(&categoryResult.ID, &categoryResult.Description,
		&categoryResult.Priority, &categoryResult.Status)

	if err != nil {
		fmt.Print(err)
		return model.Category{}, err
	}

	return categoryResult, nil
}

func (cr *CategoryRepository) ChangeStatus(id int) (model.Category, error) {

	var category model.Category

	query, err := cr.connection.Prepare("update category set status = not status where id = $1 returning *")

	if err != nil {
		fmt.Print(err)
		return model.Category{}, err
	}

	err = query.QueryRow(id).Scan(&category.ID, &category.Description, &category.Priority, &category.Status)

	if err != nil {
		fmt.Print(err)
		return model.Category{}, err
	}

	return category, nil

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

	return "Delatado com sucesso", nil
}

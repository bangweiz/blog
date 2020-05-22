package models

import (
	"database/sql"
	"time"
)

type Category struct {
	BaseModel
	Title string `json:"title"`
}

func NewCategory(title string) (ok bool, res *Category) {
	var id int
	queryString := "SELECT * FROM categories WHERE title = ?"
	err := DB.QueryRow(queryString, title).Scan(&id)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, &Category{}
		}
		queryString = "INSERT INTO categories (title, created_on, modified_on) VALUES (?, ?, ?)"
		res, err := DB.Exec(queryString, title, time.Now(), time.Now())
		if err != nil {
			return false, &Category{}
		}
		id, err := res.LastInsertId()
		if err != nil {
			return false, &Category{}
		}
		c := &Category{
			BaseModel{
				ID: int(id),
			},
			title,
		}
		return true, c
	}
	return false, &Category{}
}

func DeleteCategory(id string) bool {
	queryString := "DELETE FROM categories WHERE id = ?"
	_, err := DB.Exec(queryString, id)
	if err != nil {
		return false
	}
	return true
}

func FetchCategories() []*Category {
	queryString := "SELECT * FROM categories"
	rows, err := DB.Query(queryString)
	if err != nil {
		return nil
	}
	categories := make([]*Category, 0)
	for rows.Next() {
		var id int
		var createdOn, modifiedOn, title string
		err = rows.Scan(&id, &title, &createdOn, &modifiedOn)
		if err != nil {
			return nil
		}
		category := &Category{
			BaseModel{
				id,
				createdOn,
				modifiedOn,
			},
			title,
		}
		categories = append(categories, category)
	}
	return categories
}
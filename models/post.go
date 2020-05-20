package models

import (
	"time"
)

type Post struct {
	BaseModel
	Title string `json:"title"`
	Content string `json:"content"`
	User *User `json:"user_id"`
	Category *Category `json:"category_id"`
}

func SavePost(title string, content string, userId string, categoryId string) bool {
	queryString := "INSERT INTO posts(title, content, user_id, category_id, created_on, modified_on) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := DB.Exec(queryString, title, content, userId, categoryId, time.Now(), time.Now())
	if err != nil {
		return false
	}
	return true
}

func FetchPosts() []*Post {
	queryString := "SELECT posts.id AS postId, posts.title AS postsTitle, content, posts.created_on, posts.modified_on, user_id, users.username, category_id, categories.title AS categoryTitle FROM posts JOIN categories ON category_id = categories.id JOIN users ON user_id = users.id"
	rows, err := DB.Query(queryString)
	if err != nil {
		return nil
	}
	posts := make([]*Post, 0)
	for rows.Next() {
		var id, userId, categoryId int
		var title, content, username, categoryTitle, createdOn, modifiedOn string
		err = rows.Scan(&id, &title, &content, &createdOn, &modifiedOn, &userId, &username, &categoryId, &categoryTitle)
		if err != nil {
			return nil
		}
		p := &Post{
			BaseModel{
				id,
				createdOn,
				modifiedOn,
			},
			title,
			content,
			&User{
				BaseModel: BaseModel{
					ID: userId,
				},
				Username: username,
			},
			&Category{
				BaseModel{
					ID: categoryId,
				},
				categoryTitle,
			},
		}
		posts = append(posts, p)
	}
	return posts
}
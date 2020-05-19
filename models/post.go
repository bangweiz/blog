package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	CreateOn primitive.Timestamp `bson:"created_on,omitempty"`
	ModifiedOn primitive.Timestamp `bson:"modified_on,omitempty"`
	Title string `bson:"title,omitempty"`
	Content string `bson:"content,omitempty"`
	User primitive.ObjectID `bson:"user,omitempty"`
	Category primitive.ObjectID `bson:"category,omitempty"`
}

func SavePost(title string, content string, user string, category string) bool {
	userID, _ := primitive.ObjectIDFromHex(user)
	categoryID, _ := primitive.ObjectIDFromHex(category)
	post := Post{
		CreateOn: primitive.Timestamp{T:uint32(time.Now().Unix())},
		ModifiedOn: primitive.Timestamp{T:uint32(time.Now().Unix())},
		Title: title,
		Content: content,
		User: userID,
		Category: categoryID,
	}
	_, err := DB.PostsCollection.InsertOne(CTX, post)
	if err != nil {
		return false
	}
	return true
}

func FetchPosts() (posts []Post) {
	cursor, err := DB.PostsCollection.Find(CTX, bson.M{})
	if err != nil {
		return
	}
	if err = cursor.All(CTX, &posts); err != nil {
		return
	}
	return
}
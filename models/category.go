package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	CreateOn primitive.Timestamp `bson:"created_on,omitempty"`
	ModifiedOn primitive.Timestamp `bson:"modified_on,omitempty"`
	Title string `json:"title,omitempty"`
}

func NewCategory(title string) (ok bool, res *Category) {
	cursor, err := DB.CategoriesCollection.Find(CTX, bson.M{"title": title});
	if err != nil || cursor.Next(CTX) {
		return
	}
	category := Category{
		CreateOn: primitive.Timestamp{T:uint32(time.Now().Unix())},
		ModifiedOn: primitive.Timestamp{T:uint32(time.Now().Unix())},
		Title: title,
	}
	_, err = DB.CategoriesCollection.InsertOne(CTX, category)
	if err != nil {
		return false, &Category{}
	}
	return true, &category
}

func DeleteCategory(id string) bool {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false
	}
	_, err = DB.CategoriesCollection.DeleteOne(CTX, bson.M{"_id": primitiveID})
	if err != nil {
		return false
	}
	return true
}

func FetchCategories() []Category {
	var categories []Category
	cursor, err := DB.CategoriesCollection.Find(CTX, bson.M{})
	if err != nil {
		return categories
	}
	if err = cursor.All(CTX, &categories); err != nil {
		return categories
	}
	return categories
}
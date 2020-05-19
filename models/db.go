package models

import (
	"context"
	"log"
	"time"

	"github.com/bangweiz/blog/pkg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB *Database
	CTX context.Context
)

type Database struct {
	Client *mongo.Client
	DB *mongo.Database
	UsersCollection *mongo.Collection
	CategoriesCollection *mongo.Collection
	PostsCollection *mongo.Collection
}


func init() {
	clientOptions := options.Client().ApplyURI(pkg.URI)
	clientOptions.SetMaxPoolSize(50)

	CTX, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	client, err := mongo.Connect(CTX, clientOptions)
	if err != nil {
		log.Fatalln(err)
	}

	//defer client.Disconnect(CTX)

	database := client.Database("blog")
	usersCollection := database.Collection("users")
	categoriesCollection := database.Collection("categories")
	postsCollection := database.Collection("posts")

	DB = &Database{
		Client: client,
		DB: database,
		UsersCollection: usersCollection,
		CategoriesCollection: categoriesCollection,
		PostsCollection: postsCollection,
	}
}

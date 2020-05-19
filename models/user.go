package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	CreateOn primitive.Timestamp `bson:"created_on,omitempty"`
	ModifiedOn primitive.Timestamp `bson:"modified_on,omitempty"`
	Username string `bson:"username,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
}

func UserRegister(username string, email string, pwd string) (ok bool) {
	cursor, err := DB.UsersCollection.Find(CTX, bson.M{"email": email})
	if err != nil {
		return
	}
	if cursor.Next(CTX) {
		return
	}
	user := User{
		CreateOn: primitive.Timestamp{T:uint32(time.Now().Unix())},
		ModifiedOn: primitive.Timestamp{T:uint32(time.Now().Unix())},
		Username: username,
		Email: email,
		Password: pwd,
	}
	_, err = DB.UsersCollection.InsertOne(CTX, user)
	if err != nil {
		return
	}
	return true
}

func FindUser(email string, pwd string) (bool, User) {
	var user User
	if err := DB.UsersCollection.FindOne(CTX, bson.M{"email": email}).Decode(&user); err != nil {
		return false, User{}
	}
	if user.Password != pwd {
		return false, User{}
	}
	return true, user
}
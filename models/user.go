package models

import (
	"context"

	"github.com/MayamkSaxena03/Accuknox/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserData struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

func GetUserCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "users")
}

func (u *UserData) ValidateSignupBody() bool {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return false
	}
	return true
}

func (u *UserData) ValidateLoginBody() bool {
	if u.Email == "" || u.Password == "" {
		return false
	}
	return true
}

func CheckUserExists(ctx context.Context, email string) (bool, error) {
	userCollection := GetUserCollection()
	var user UserData
	err := userCollection.FindOne(ctx, UserData{Email: email}).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		return false, err
	}

	if user.Email == email {
		return true, nil
	}

	return false, nil
}

package models

import (
	"github.com/MayamkSaxena03/Accuknox/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteData struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"userid,omitempty" bson:"userid,omitempty"`
	Note   string             `json:"note,omitempty" bson:"note,omitempty"`
}

type NoteDataRequest struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Note string             `json:"note,omitempty" bson:"note,omitempty"`
	SID  string             `json:"sid,omitempty" bson:"sid,omitempty"`
}

type AllNotes struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Note string             `json:"note,omitempty" bson:"note,omitempty"`
}

func GetNoteCollection() *mongo.Collection {
	return database.OpenCollection(database.Client, "notes")
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Dish        string             `json:"dish"`
	Fat         float64            `json:"fat"`
	Ingredients string
	Calories    string
}

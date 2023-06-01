package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Movie   string             `json:"movie" bson:"movie"`
	Watched bool               `json:"is_watched" bson:"is_watched"`
}


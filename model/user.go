package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Admin    int                `json:"admin,omitempty" bson:"admin"`
	Name     string             `json:"name,omitempty" bson:"name"`
	Password string             `json:"password,omitempty" bson:"password"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bangumi struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `json:"name,omitempty" bson:"name"`
	EpisodeNumbers int                `json:"episode_numbers,omitempty" bson:"episode_numbers"`
	EpisodeName    []string           `json:"episode_name,omitempty" bson:"episode_name"`
	StartTime      string             `json:"start_time,omitempty" bson:"start_time"`
	Discription    string             `json:"discription,omitempty" bson:"discription"`
}

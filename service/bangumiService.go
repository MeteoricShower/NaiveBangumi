package service

import (
	"NaiveBangumi/database"
	"NaiveBangumi/model"
	"go.mongodb.org/mongo-driver/bson"
)

func FindBangumi(filter bson.M) (model.Bangumi, error) {
	bagnumi := model.Bangumi{}
	result := database.BangumiCollection.FindOne(database.Ctx, filter)
	err := result.Decode(&bagnumi)
	if err != nil {
		return bagnumi, err
	}
	return bagnumi, nil
}

func InsertBangumi(bangumi model.Bangumi) error {
	_, err := database.BangumiCollection.InsertOne(database.Ctx, bson.M{
		"name":            bangumi.Name,
		"episode_numbers": bangumi.EpisodeNumbers,
		"episode_name":    bangumi.EpisodeName,
		"start_time":      bangumi.StartTime,
		"discription":     bangumi.Discription,
	})
	if err != nil {
		return err
	}
	return nil
}

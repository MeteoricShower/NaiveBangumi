package service

import (
	"NaiveBangumi/database"
	"NaiveBangumi/model"
	"go.mongodb.org/mongo-driver/bson"
)

func BangumiUpdateRequest(bangumiModifyRequest model.BangumiModifyRequest) error {
	_, err := database.BangumiUpdateRequestCollection.InsertOne(database.Ctx, bson.M{
		"name":            bangumiModifyRequest.Name,
		"episode_numbers": bangumiModifyRequest.EpisodeNumbers,
		"episode_name":    bangumiModifyRequest.EpisodeName,
		"start_time":      bangumiModifyRequest.StartTime,
		"discription":     bangumiModifyRequest.Discription,
		"sender":          bangumiModifyRequest.Sender,
	})
	if err != nil {
		return err
	}
	return nil
}

package service

import (
	"NaiveBangumi/database"
	"NaiveBangumi/model"
	"go.mongodb.org/mongo-driver/bson"
)

//找一个番剧
func FindBangumi(filter bson.M) (model.Bangumi, error) {
	bagnumi := model.Bangumi{}
	result := database.BangumiCollection.FindOne(database.Ctx, filter)
	err := result.Decode(&bagnumi)
	if err != nil {
		return bagnumi, err
	}
	return bagnumi, nil
}

//找符合条件的所有番剧
func FindBangumiAll(filter bson.M) ([]model.Bangumi, error) {
	cursor, err := database.BangumiCollection.Find(database.Ctx, filter)
	if err != nil {
		return nil, err
	}

	//查找多个文档返回一个游标
	//遍历游标一次解码一个游标
	var res []model.Bangumi

	err = cursor.All(database.Ctx, &res)
	//for cursor.Next(database.Ctx) {
	//	//定义一个文档，将单个文档解码为p
	//	var p model.Bangumi
	//	err := cursor.Decode(&p)
	//	if err != nil {
	//		return nil, err
	//	}
	//	res = append(res, p)
	//}
	return res, nil
}

//添加一个番剧
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

func UpdateBangumi(bangumi model.Bangumi) error {
	filter := bson.D{{"name", bangumi.Name}}
	update := bson.M{
		"$set": bson.M{
			"name":            bangumi.Name,
			"episode_numbers": bangumi.EpisodeNumbers,
			"episode_name":    bangumi.EpisodeName,
			"start_time":      bangumi.StartTime,
			"discription":     bangumi.Discription},
	}
	_, err := database.BangumiCollection.UpdateOne(database.Ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

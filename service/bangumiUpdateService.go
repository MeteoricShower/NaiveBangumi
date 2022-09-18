package service

import (
	"NaiveBangumi/database"
	"NaiveBangumi/model"
	"go.mongodb.org/mongo-driver/bson"
)

//把修改番剧请求存入数据库
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

//更新修改番剧接收状态
func UpdateRequest(bangumiModifyRequest model.BangumiModifyRequest) error {
	filter := bson.D{{"name", bangumiModifyRequest.Name}}
	update := bson.M{
		"$set": bson.M{
			"is_received": true},
	}
	_, err := database.BangumiCollection.UpdateOne(database.Ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

//查找所有修改番剧请求
func FindBangumiUpdateRequest(filter bson.M) ([]model.BangumiModifyRequest, error) {
	cursor, err := database.BangumiUpdateRequestCollection.Find(database.Ctx, filter)
	if err != nil {
		return nil, err
	}

	//查找多个文档返回一个游标
	//遍历游标一次解码一个游标
	var res []model.BangumiModifyRequest
	err = cursor.All(database.Ctx, &res)
	for _, bmr := range res {
		err := UpdateRequest(bmr)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

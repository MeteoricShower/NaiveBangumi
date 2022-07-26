package service

import (
	"NaiveBangumi/database"
	"NaiveBangumi/model"
	"go.mongodb.org/mongo-driver/bson"
)

//添加一个用户
func InsertUser(user model.User) error {
	_, err := database.UserCollection.InsertOne(database.Ctx, bson.M{
		"admin":    user.Admin,
		"name":     user.Name,
		"password": user.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

//查找一个用户
func FindUser(filter bson.M) (model.User, error) {
	user := model.User{}
	result := database.UserCollection.FindOne(database.Ctx, filter)
	err := result.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil

}

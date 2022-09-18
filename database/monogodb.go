package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "time"
)

var UserCollection *mongo.Collection
var BangumiCollection *mongo.Collection
var BangumiUpdateRequestCollection *mongo.Collection
var Ctx context.Context

func Init() {
	// 设置客户端选项
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接 MongoDB
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	UserCollection = client.Database("NaiveBangumi").Collection("users")
	BangumiCollection = client.Database("NaiveBangumi").Collection("bangumi")
	BangumiUpdateRequestCollection = client.Database("NaiveBangumi").Collection("bangumi_update_request")
	fmt.Println("Connected to MongoDB!")

}

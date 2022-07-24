package main

import (
	"NaiveBangumi/database"
	"NaiveBangumi/router"
)

func main() {
	database.Init()
	//service.InsertUser(model.User{Name: "aaa", Password: "7777"})
	//user, _ := service.FindUser(bson.M{"name": "aaa"})
	//println(user.Password)
	router.Run()
}

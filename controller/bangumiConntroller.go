package controller

import (
	"NaiveBangumi/model"
	"NaiveBangumi/service"
	"NaiveBangumi/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AddBangumiOne(c echo.Context) (err error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Jwt)
	admin := claims.Admin
	if admin != 1 {
		return c.JSON(http.StatusUnauthorized, util.MyInfo{Success: false, Error: "You are not admministrator"})
	}
	// Bind
	b := model.Bangumi{}
	err = c.Bind(&b)
	if err != nil {
		return
	}
	//check whether user is exist
	var filter = bson.M{"name": b.Name}
	bangumi, err := service.FindBangumi(filter)
	if bangumi.Name != "" {
		return c.JSON(http.StatusBadRequest, util.MyInfo{Success: false, Error: "Bangumi is already exist"})
	}
	err = service.InsertBangumi(b)
	if err != nil {
		return
	}
	return c.JSON(http.StatusCreated, b)
}

func FindBangumi(c echo.Context) (err error) {
	b := model.Bangumi{}
	err = c.Bind(&b)
	if err != nil {
		return
	}
	//根据名字找番剧
	filter := bson.M{"name": bson.M{"$regex": primitive.Regex{Pattern: ".*" + b.Name + ".*", Options: "i"}}}
	//filter := bson.M{"name": bson.M{"$regex": "/" + b.Name + "/"}}
	bList, err := service.FindBangumiAll(filter)
	if err != nil {
		return
	}
	if len(bList) == 0 {
		return c.JSON(http.StatusNotFound, util.MyInfo{Success: false, Error: "Bangumi is not exist"})
	}
	return c.JSON(http.StatusOK, bList)
}

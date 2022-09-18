package controller

import (
	"NaiveBangumi/model"
	"NaiveBangumi/service"
	"NaiveBangumi/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

//发送修改番剧请求
func SendBangumiUpdateRequest(c echo.Context) (err error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Jwt)
	admin := claims.Admin
	if admin == 1 {
		return c.JSON(http.StatusUnauthorized, util.MyInfo{Success: false, Error: "You are a admministrator"})
	}
	// Bind
	bmr := model.BangumiModifyRequest{}
	err = c.Bind(&bmr)
	if err != nil {
		return
	}
	err = service.BangumiUpdateRequest(bmr)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, util.MyInfo{Success: true, Data: "BangumiModifyRequest is sent"})
}

//管理员处理修改番剧请求
func HandleBangumiUpdateRequest(c echo.Context) (err error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Jwt)
	admin := claims.Admin
	if admin == 1 {
		return c.JSON(http.StatusUnauthorized, util.MyInfo{Success: false, Error: "You are not a admministrator"})
	}
	filter := bson.M{"is_received": false}
	burList, err := service.FindBangumiUpdateRequest(filter)
	for _, bmr := range burList {
		b := model.Bangumi{}
		b.Name = bmr.Name
		b.StartTime = bmr.StartTime
		b.Discription = bmr.Discription
		b.EpisodeName = bmr.EpisodeName
		b.EpisodeNumbers = bmr.EpisodeNumbers
		err := service.UpdateBangumi(b)
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, util.MyInfo{Success: true, Data: "BangumiModifyRequest is handled"})
}

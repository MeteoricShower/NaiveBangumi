package controller

import (
	"NaiveBangumi/model"
	"NaiveBangumi/service"
	"NaiveBangumi/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

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

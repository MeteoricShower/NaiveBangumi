package controller

import (
	"NaiveBangumi/model"
	"NaiveBangumi/service"
	"NaiveBangumi/util"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func Regist(c echo.Context) (err error) {
	// Bind
	u := model.User{}
	err = c.Bind(&u)
	if err != nil {
		return
	}

	// Validate
	if u.Password == "" || u.Name == "" {
		return c.JSON(http.StatusBadRequest, util.MyInfo{Success: false, Error: "invalid password or username"})
	}

	//check whether user is exist
	var filter = bson.M{"name": u.Name}
	user, err := service.FindUser(filter)
	if user.Name != "" {
		return c.JSON(http.StatusBadRequest, util.MyInfo{Success: false, Error: "User is already exist"})
	}

	// Save user
	u.Admin = 0 //普通用户权限为0
	err = service.InsertUser(u)
	if err != nil {
		return
	}
	return c.JSON(http.StatusCreated, u)
}

func Login(c echo.Context) (err error) {
	// Bind
	u := model.User{}
	err = c.Bind(&u)
	if err != nil {
		return
	}

	//check username and password
	filter := bson.M{"name": u.Name, "password": u.Password}
	user, err := service.FindUser(filter)
	if user.Name == "" {
		return c.JSON(http.StatusBadRequest, util.MyInfo{Success: false, Error: "Incorrect Username or Password"})
	}

	//-----
	// JWT
	//-----

	data := model.Jwt{
		Name:  user.Name,
		Admin: user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}
	jwts := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, _ := jwts.SignedString([]byte(`secret`))
	util.Set(u.Name, token)
	return c.JSON(http.StatusOK, util.MyInfo{Success: true, Data: token})
}

func Welcome(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.Jwt)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!"+"admin:"+strconv.Itoa(claims.Admin))
}

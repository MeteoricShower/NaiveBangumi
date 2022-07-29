package router

import (
	"NaiveBangumi/controller"
	"NaiveBangumi/model"
	_ "github.com/gorilla/sessions"
	"github.com/labstack/echo"
	_ "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

func Run() {
	e := echo.New()

	e.POST("/regist", controller.Regist)
	e.POST("/login", controller.Login)

	// Restricted group
	r := e.Group("/bangumi")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &(model.Jwt{}),
		SigningKey: []byte("secret"),
	}
	r.Use(middleware.JWTWithConfig(config))
	r.GET("", controller.Welcome)
	r.GET("/find", controller.FindBangumi)
	r.POST("/add/one", controller.AddBangumiOne)
	r.POST("/update", controller.UpdateBangumi)
	e.Logger.Fatal(e.Start(":1323"))
}

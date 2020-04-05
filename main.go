package main

import (
	"ssm/handler"
	"ssm/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.GateKeeper)
	e.GET("/:name", handler.ServeHomePath)
	e.GET("/json", handler.ServeHomePathReturnJson)
	e.POST("/login", handler.ServeHomePathReceivePostBody)
	e.Logger.Fatal(e.Start(":1323"))
}

package handler

import (
	"fmt"
	"net/http"
	"ssm/model"

	"github.com/labstack/echo/v4"
)

// ServeHomePath to render homepage
func ServeHomePath(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello : "+name)
}

type simpleJsonResponse struct {
	Username string
	Point    float64
}

// ServeHomePathReturnJson to render homepage
func ServeHomePathReturnJson(c echo.Context) error {
	jsonObj := simpleJsonResponse{
		"aorjoa", 10.2,
	}
	return c.JSON(http.StatusOK, jsonObj)
}

type loginRequest struct {
	Username string
	Password string
}

// ServeHomePathReceivePostBody to receive post body
func ServeHomePathReceivePostBody(uls model.UserLoginLogDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody loginRequest
		err := c.Bind(&reqBody)
		if err != nil {
			fmt.Println("something wrong")
		}

		userLoginLog := &model.UserLoginLog{Username: reqBody.Username, Password: reqBody.Password}
		err = uls.Create(userLoginLog)
		if err != nil {
			fmt.Println("error when creating")
		}
		return c.JSON(http.StatusOK, reqBody)
	}
}

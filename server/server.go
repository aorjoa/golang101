package server

import (
	"fmt"
	"log"
	"net/http"
	"ssm/handler"
	"ssm/middleware"
	"ssm/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Server is a struct for abstraction server
type Server struct {
	Echo           *echo.Echo
	DB             *gorm.DB
	UserLoginLogDB model.UserLoginLogDB
}

// ServerRoute is a method to serve path
func (s *Server) ServerRoute() {
	e := s.Echo
	e.Use(middleware.GateKeeper)
	e.GET("/:name", handler.ServeHomePath)
	e.GET("/json", handler.ServeHomePathReturnJson)
	e.POST("/login", handler.ServeHomePathReceivePostBody(s.UserLoginLogDB))
	e.POST("/callback", func(c echo.Context) error {
		// Line bot
		bot, err := linebot.New("1bcd268701bde76a7adfeb34af402b49", "")
		/////////////
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.NoContent(http.StatusBadRequest)
			} else {
				c.NoContent(http.StatusInternalServerError)
			}
			return err
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					fmt.Println(message.Text)

					if message.Text == "Test" {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("BoooYahhhh!")).Do()
						continue
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
		return c.NoContent(http.StatusOK)
	})
}

// ServeHTTP should serve an enpoint
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Echo.ServeHTTP(w, r)
}

// Start server at specified port
func (s *Server) Start(address string) error {
	return s.Echo.Start(address)
}

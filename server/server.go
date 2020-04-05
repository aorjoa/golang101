package server

import (
	"net/http"
	"ssm/handler"
	"ssm/middleware"
	"ssm/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Server is a struct for abstraction server
type Server struct {
	Echo                    *echo.Echo
	DB                      *gorm.DB
	UserLoginLogDB			model.UserLoginLogDB
}

// ServerRoute is a method to serve path
func (s *Server) ServerRoute() {
	e := s.Echo
	e.Use(middleware.GateKeeper)
	e.GET("/:name", handler.ServeHomePath)
	e.GET("/json", handler.ServeHomePathReturnJson)
	e.POST("/login", handler.ServeHomePathReceivePostBody(s.UserLoginLogDB))
}

// ServeHTTP should serve an enpoint
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Echo.ServeHTTP(w, r)
}

// Start server at specified port
func (s *Server) Start(address string) error {
	return s.Echo.Start(address)
}
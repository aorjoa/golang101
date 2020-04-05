package main

import (
	"ssm/model"
	"ssm/server"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&model.Product{}, &model.UserLoginLog{})

	e := echo.New()
	server := server.Server{
		Echo:           e,
		DB:             db,
		UserLoginLogDB: model.NewUserLoginLogDB(db),
	}
	server.ServerRoute()
	server.Start(":1323")
}

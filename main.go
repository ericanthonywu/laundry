package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"laundry/Config"
	"laundry/Middleware"
	"laundry/Model/Database"
	"laundry/Route"
	"laundry/Utils"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	// run migration
	err := Config.Db().AutoMigrate(
		&Database.User{},
		&Database.Laundry{},
		&Database.UserOtpRequest{},
	)
	if err != nil {
		panic(err)
	}

	Middleware.Init(e)
	Route.Init(e)

	e.Logger.Fatal(e.Start(":" + Utils.GetEnv("PORT")))
}

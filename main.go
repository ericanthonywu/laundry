package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"laundry/Lib"
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

	Lib.InitAll()

	// run migration
	err := Lib.DB.AutoMigrate(
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

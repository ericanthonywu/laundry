package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"laundry/Middleware"
	"laundry/Route"
	"laundry/Utils"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err)
	}

	Middleware.Init(e)
	Route.Init(e)

	e.Logger.Fatal(e.Start(":" + Utils.GetEnvDefault("PORT", "8000")))
}

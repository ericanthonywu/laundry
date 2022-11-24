package Route

import (
	"github.com/labstack/echo/v4"
	"laundry/Controller/User"
	"laundry/Middleware"
)

func UserRoute(e *echo.Echo) {
	api := e.Group("/user")

	// for authentication route
	api.POST("/request-otp", User.RequestOTP)
	api.POST("/verify-otp", User.VerifyOTP)

	// protected route
	api.GET("/getUserData", User.GetProfile, Middleware.UserMiddleware)
}

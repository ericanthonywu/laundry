package Route

import (
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	go IndexRoute(e)
	go UserRoute(e)
}

func IndexRoute(e *echo.Echo) {

}

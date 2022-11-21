package Route

import (
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	IndexRoute(e)
	UserRoute(e)
}

func IndexRoute(e *echo.Echo) {

}

package User

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"laundry/Utils"
)

func GetProfile(c echo.Context) error {
	id, role := Utils.GetJwtClaims(c)
	fmt.Println(id + " - " + role)
	return nil
}

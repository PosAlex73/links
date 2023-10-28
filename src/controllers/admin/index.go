package admin

import "github.com/labstack/echo/v4"

func Index(context echo.Context) error {
	return context.JSON(200, "123")
}

package router

import (
	"github.com/labstack/echo/v4"
	"links/src/controllers/admin"
)

func BindRoutes(server *echo.Echo) {
	server.GET("/", admin.Index)
}

package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"links/src/application"
)

func main() {
	app := application.Application{
		Server: echo.New(),
	}

	app.Run()
}

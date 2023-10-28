package application

import (
	"github.com/labstack/echo/v4"
	"links/src/router"
	"os"
	"strings"
)

const MainConfig = "./config/main.yaml"

type Application struct {
	Server *echo.Echo
	config map[string]string
}

func (app *Application) Run() {
	app.init()

	host := app.config["host"]
	port := app.config["port"]

	app.Server.Logger.Fatal(app.Server.Start(host + ":" + port))
}

func (app *Application) init() {
	router.BindRoutes(app.Server)
	app.config = readConfig()
}

func readConfig() map[string]string {
	config := make(map[string]string)

	configRawContent, err := os.ReadFile(MainConfig)
	configContentString := string(configRawContent)
	parseConfigArray := strings.Split(configContentString, "\n")

	if len(parseConfigArray) > 0 {
		for _, settingRow := range parseConfigArray {
			settings := strings.Split(settingRow, " ")
			settings[0] = strings.Trim(settings[0], ": \n")
			settings[1] = strings.Trim(settings[1], " \n")
			config[settings[0]] = settings[1]
		}
	}

	if err != nil {
		panic("Не удалось прочить основной файл конфигурации")
	}

	return config
}

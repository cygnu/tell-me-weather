package main

import (
	"fmt"
	"github.com/cygnu/tell-me-weather/config"
	"github.com/cygnu/tell-me-weather/openweather"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", handler)

	start := e.Start(fmt.Sprintf(":%v", config.Config.Port))
	e.Logger.Fatal(start)
}

func handler(c echo.Context) error {
	apiClient, _ := openweather.NewAPIClient(openweather.BaseURL, config.Config.ApiKey)
	apiClient.GetForecast("Tokyo,jp")

	return nil
}
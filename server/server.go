package server

import (
	"fmt"
	"github.com/cygnu/tell-me-weather/config"
	"github.com/cygnu/tell-me-weather/openweather"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	apiClient, _ := openweather.NewAPIClient(openweather.BaseURL, config.Config.ApiKey)
	//apiClient.MakeRequest("forecast", "Tokyo,jp")
	apiClient.GetForecast("Tokyo,jp")
}

func StartWebServer() error {
	http.HandleFunc("/", handler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}

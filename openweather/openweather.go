package openweather

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cygnu/tell-me-weather/config"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var BaseURL = "https://api.openweathermap.org/data/2.5"

type APIClient struct {
	baseURL    *url.URL
	key        string
	httpClient *http.Client
}

func NewAPIClient(baseURL, key string) (*APIClient, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	apiClient := &APIClient{parsedURL, key, &http.Client{}}
	return apiClient, nil
}

func (api *APIClient) doRequest(req *http.Request) (body []byte, err error) {
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (api *APIClient) MakeRequest(context, city string) (*http.Request, error) {
	values := url.Values{}
	values.Set("q", city)
	values.Set("appid", config.Config.ApiKey)

	req := fmt.Sprintf("%s/%s?%s", BaseURL, context, values.Encode())
	log.Printf("action=makeRequest url=%s", req)

	return http.NewRequest("GET", req, nil)
}

type Forecast struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []List `json:"list"`
}

type List struct {
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	DtTxt   string    `json:"dt_txt"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func (api *APIClient) GetForecast(city string) (*Forecast, error) {
	reqURL, err := api.MakeRequest("forecast", city)
	if err != nil {
		return nil, err
	}

	resp, err := api.doRequest(reqURL)
	log.Printf("resp=%s\n", string(resp))
	if err != nil {
		log.Printf("action=GetForecast err=%s", err.Error())
		return nil, err
	}

	var forecast Forecast
	err = json.Unmarshal(resp, &forecast)
	if err != nil {
		log.Printf("action=GetForecast err=%s", err.Error())
		return nil, err
	}
	return &forecast, nil
}

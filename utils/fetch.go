package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"alexleyoung/sun/types"
)

func GetForecast(apiKey string, location string, days int) types.ForecastResponse {
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%d", apiKey, location, days))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		var error types.Error;
		err := json.Unmarshal(body, &error)
		if err != nil {
			panic(err)
		} else {
			panic(error.E.Message)
		}
	}

	var forecastResponse types.ForecastResponse

	err = json.Unmarshal(body, &forecastResponse)
	if err != nil {
		panic(err)
	}

	return forecastResponse
}

func GetAlerts(apiKey string, location string) types.AlertsResponse {
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/alerts.json?key=%s&q=%s", apiKey, location))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	if res.StatusCode != 200 {
		var error types.Error;
		err := json.Unmarshal(body, &error)
		if err != nil {
			panic(err)
		} else {
			panic(error.E.Message)
		}
	}

	var alertResponse types.AlertsResponse

	err = json.Unmarshal(body, &alertResponse)
	if err != nil {
		panic(err)
	}

	return alertResponse
}
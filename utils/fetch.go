package utils

import (
	"alexleyoung/sun/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	var forecast types.ForecastResponse

	err = json.Unmarshal(body, &forecast)
	if err != nil {
		panic(err)
	}

	return forecast
}

func GetAlerts(apiKey string, location string) {
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

	fmt.Println(string(body))

	// var alerts types.Alerts

	// err = json.Unmarshal(body, &alerts)
	// if err != nil {
	// 	panic(err)
	// }

	// return alerts
}
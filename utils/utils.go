package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetWeatherRestOfDay(apiKey string, location string) []byte {
	res, err := http.Get(fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s", apiKey, location))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(fmt.Sprintf("Request failed with status code %d", res.StatusCode))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return body
}
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"alexleyoung/sun/types"
)

func GetForecast(apiKey string, location string, days int) []byte {
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

	return body
}
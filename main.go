package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Location struct {
	Zip     string `json:"zip"`
	Name    string `json:"name"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
	Country string `json:"country"`
}

type Weather struct {
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp          float32 `json:"temp"`
		FeelsLike     float32 `json:"feels_like"`
		TempMin       float32 `json:"temp_min"`
		TempMax       float32 `json:"temp_max"`
		Pressure      int     `json:"pressure"`
		Humidity      int     `json:"humidity"`
		SeaLevel      int     `json:"sea_level"`
		GrndLevel     int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float32 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float32 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int `json:"timezone"`
	ID       int `json:"id"`
	Name     string `json:"name"`
	Cod      int `json:"cod"`
}

func setAPIKey(key string) {
	os.Setenv("API_KEY", key)
}

func setZip(zip string) {
	os.Setenv("ZIP", zip)
}

func setCountry(country string) {
	os.Setenv("COUNTRY", country)
}

func main() {
	setZip("50014")
	setCountry("US")
	setAPIKey("23a21d0da3a7dcc31eb1fff683324c8e")
	units := "imperial"

	API_KEY := os.Getenv("API_KEY")
	ZIP := os.Getenv("ZIP")
	COUNTRY := os.Getenv("COUNTRY")

	res, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/zip?zip=%s,%s&appid=%s", ZIP, COUNTRY, API_KEY))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var location Location;
	err = json.Unmarshal(body, &location)

	if err != nil {
		panic(err)
	}

	lat := location.Lat
	lon := location.Lon

	res, err = http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&units=%s&appid=%s", lat, lon, units, API_KEY))

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)

	if err != nil {
		panic(err)
	}

	fmt.Println(weather.Main.Temp)
}
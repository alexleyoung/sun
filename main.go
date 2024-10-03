package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// body looks like this: {"zip":"50014","name":"Story County","lat":42.0486,"lon":-93.6945,"country":"US"}
type Location struct {
	Zip     string `json:"zip"`
	Name    string `json:"name"`
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
	Country string `json:"country"`
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

	res, err = http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, API_KEY))

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
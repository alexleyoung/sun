package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http://api.openweathermap.org/geo/1.0/zip?zip={zip code},{country code}&appid={API key}
	res, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/geo/1.0/zip?zip=%s,%s&appid=%s", os.Getenv("ZIP_CODE"), os.Getenv("COUNTRY_CODE"), os.Getenv("API_KEY")))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(fmt.Sprintf("status code: %d", res.StatusCode))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var apiKey = "your_api_key"
var cityId = "city_name" // {city name},{state code},{country code}

var URL string = "https://api.openweathermap.org/data/2.5/weather?q=" + cityId + "&appid=" + apiKey

type Weather struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Temp_min float64 `json:"temp_min"`
		Temp_max float64 `json:"temp_max"`
	} `json:"main"`
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func kelvinToCelsius(k float64) float64 {
	c := k - 273.15
	return c
}

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal()
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	Weather1 := Weather{}

	err = json.Unmarshal(body, &Weather1)
	if err != nil {
		fmt.Println("Can not unmarshal JSON")
		log.Fatal()
	}

	fmt.Println("Id, City:", Weather1.Id, ",", Weather1.Name)
	fmt.Printf("Temp:%.1f°\n", kelvinToCelsius(Weather1.Main.Temp))
	fmt.Printf("Min_Temp:%.1f°\n", kelvinToCelsius(Weather1.Main.Temp_min))
	fmt.Printf("Max_Temp:%.1f°\n", kelvinToCelsius(Weather1.Main.Temp_max))

}

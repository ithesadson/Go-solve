package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var apiKey = "your_api_key"
var cityId = "city_id"
var URL string = "http://api.openweathermap.org/data/2.5/forecast?id=" + cityId + "&appid=" + apiKey

type Weather struct {
	List []struct {
		Main struct {
			Temp    float64 `json:"temp"`
			Tempmin float64 `json:"temp_min"`
			Tempmax float64 `json:"temp_max"`
		} `json:"main"`
	} `json:"list"`
	City struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"city"`
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

	fmt.Println("Id, City:", Weather1.City.Id, ",", Weather1.City.Name)
	fmt.Printf("Temp:%.2f°\n", kelvinToCelsius(Weather1.List[0].Main.Temp))
	fmt.Printf("Min_Temp:%.2f°\n", kelvinToCelsius(Weather1.List[0].Main.Tempmin))
	fmt.Printf("Max_Temp:%.2f°\n", kelvinToCelsius(Weather1.List[0].Main.Tempmax))
}

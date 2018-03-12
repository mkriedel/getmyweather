package main

import (
	"flag"

	//"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/mkriedel/getmyweather/gmw"
	"fmt"
)

var (
	searchquery string
)

// Use the '-city' flag to search for a city.
func main() {
	flag.StringVar(&searchquery, "city", "Auckland", "The name of the city to search for.")
	flag.Parse()
	cities, _ := gmw.GetCityInfo(searchquery)
	//debugWeather(cities[0].Woeid)
	for _, city := range cities {
		weather := gmw.GetWeather(city.Woeid)
		fmt.Printf("Weather for %s, %s (%d)\n", weather.Title, weather.Parent.Title, city.Woeid)
		spew.Dump(weather)
	}
	//fmt.Printf("%+v\n", weather)
	//fmt.Println(weather.Latt_Long.Latitude)
	//fmt.Println(weather.Time, weather.Sun_Rise, weather.Sun_Set)
	//fmt.Println(createTempList(weather.Consolidated_Weather))
	//fmt.Println(AverageTemp(weather.Consolidated_Weather))
	//fmt.Println(MedianTemp(weather.Consolidated_Weather))
}

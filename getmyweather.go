package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//"sort"

	"github.com/davecgh/go-spew/spew"
	//"github.com/montanaflynn/stats"
	//"github.com/tidwall/gjson"
)

var (
	searchquery string
)

func fetchInfo(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}
	return info, nil
}

func getWeather(woeid int) Weather {
	weather := Weather{}
	apiUrl := fmt.Sprintf("https://www.metaweather.com/api/location/%d/", woeid)
	resp, _ := fetchInfo(apiUrl)
	json.Unmarshal(resp, &weather)
	return weather
}

func getCityInfo(city string) ([]CityInfo, error) {
	cityEncoded := url.PathEscape(city)
	cities := []CityInfo{}
	apiUrl := fmt.Sprintf("https://www.metaweather.com/api/location/search/?query=%s", cityEncoded)
	info, _ := fetchInfo(apiUrl)
	json.Unmarshal(info, &cities)
	return cities, nil
}

func debugWeather(woeid int) {
	var info interface{}
	apiUrl := fmt.Sprintf("https://www.metaweather.com/api/location/%d/", woeid)
	resp, _ := fetchInfo(apiUrl)
	json.Unmarshal(resp, &info)
	prettyInfo, _ := json.MarshalIndent(info, "", "  ")
	fmt.Printf("%s\n", prettyInfo)
}

//func createTempList(weather []consolidatedWeather) []float64 {
//	var temps []float64
//	for i := 0; i < len(weather); i++ {
//		temps = append(temps, weather[i].The_Temp)
//	}
//	sort.Float64s(temps)
//	return temps
//}
//
//func AverageTemp(weather []consolidatedWeather) float64 {
//	meanTemps, _ := stats.Mean(createTempList(weather))
//	return meanTemps
//}
//
//func MedianTemp(weather []consolidatedWeather) float64 {
//	medianTemps, _ := stats.Median(createTempList(weather))
//	return medianTemps
//}

func main() {
	flag.StringVar(&searchquery, "city", "Auckland", "The name of the city to search for.")
	flag.Parse()
	cities, _ := getCityInfo(searchquery)
	//debugWeather(cities[0].Woeid)
	weather := getWeather(cities[0].Woeid)
	spew.Dump(weather)
	//fmt.Printf("%+v\n", weather)
	//fmt.Println(weather.Latt_Long.Latitude)
	//fmt.Println(weather.Time, weather.Sun_Rise, weather.Sun_Set)
	//fmt.Println(createTempList(weather.Consolidated_Weather))
	//fmt.Println(AverageTemp(weather.Consolidated_Weather))
	//fmt.Println(MedianTemp(weather.Consolidated_Weather))
}

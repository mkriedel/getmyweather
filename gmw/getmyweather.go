// getmyweather gets weather for you
package gmw

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//"sort"
	//"github.com/montanaflynn/stats"
	//"github.com/tidwall/gjson"
)

func fetchInfo(url string) ([]byte, error) {
	// Generic getter for HTML requests
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

// GetWeather gets weather
func GetWeather(woeid int) Weather {
	weather := Weather{}
	apiUrl := fmt.Sprintf("https://www.metaweather.com/api/location/%d/", woeid)
	resp, _ := fetchInfo(apiUrl)
	json.Unmarshal(resp, &weather)
	return weather
}

// GetCityInfo searches for a city from the -city command line flag and returns
// a Location{}.
func GetCityInfo(city string) ([]Location, error) {
	cityEncoded := url.PathEscape(city)
	cities := []Location{}
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

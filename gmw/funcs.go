// getmyweather gets weather for you
package gmw

import (
	"encoding/json"
	"fmt"
	"github.com/montanaflynn/stats"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func fetchInfo(url string) (info []byte, err error) {
	// Generic getter for HTML requests
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	info, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}
	return
}

// GetWeather gets weather
func GetWeather(woeid int) (weather  Weather) {
	weather = Weather{}
	apiUrl := fmt.Sprintf("https://www.metaweather.com/api/location/%d/", woeid)
	resp, _ := fetchInfo(apiUrl)
	json.Unmarshal(resp, &weather)
	return
}

// GetCityInfo searches for a city from the -city command line flag and returns
// a Location{}.
func GetCityInfo(city string) (cities []Location, err error) {
	cityEncoded := url.PathEscape(city)
	cities = []Location{}
	apiUrl := fmt.Sprintf("https://www.metaweather.com/api/location/search/?query=%s", cityEncoded)
	info, _ := fetchInfo(apiUrl)
	json.Unmarshal(info, &cities)
	return
}

func debugWeather(woeid int) {
	var info interface{}
	apiUrl := fmt.Sprintf("https://www.metaweather.com/api/location/%d/", woeid)
	resp, _ := fetchInfo(apiUrl)
	json.Unmarshal(resp, &info)
	prettyInfo, _ := json.MarshalIndent(info, "", "  ")
	fmt.Printf("%s\n", prettyInfo)
}

// UnmarshalJSON interface loads the reported temperatures into the Temp
// type during a JSON unmarshaling.
func (temps *Temp) UnmarshalJSON(b []byte) (err error) {
	temps.TempC, _ = strconv.ParseFloat(string(b), 64)
	temps.TempF = (temps.TempC * 1.8) + 32

	return
}

// UnmarshalJSON takes a comma-seperated string ("Lat,Long") into a LatLong type.
func (location *LatLong) UnmarshalJSON(b []byte) (err error) {
	location.Latitude = 0
	location.Longitude = 0
	unquoted := strings.Replace(string(b), "\"", "", -1)
	coords := strings.Split(unquoted, ",")
	if len(coords) == 2 {
		// Only set latitude and longitude if we have the coordinates.
		latitude, _ := strconv.ParseFloat(coords[0], 64)
		longitude, _ := strconv.ParseFloat(coords[1], 64)
		location.Latitude = latitude
		location.Longitude = longitude
	}

	return
}

func (location *Location) GetMapURL() (url string, err error) {
	url = fmt.Sprintf("https://www.google.com/maps?q=%v,%v", location.Latt_Long.Latitude, location.Latt_Long.Longitude)
	return
}

func (weather *Weather) createTempList() (temps []float64) {
	for i := 0; i < len(weather.Consolidated_Weather); i++ {
		temps = append(temps, weather.Consolidated_Weather[i].The_Temp.TempC)
	}
	sort.Float64s(temps)
	return
}

func (weather *Weather) AverageTemp() (meanTemps float64) {
	meanTemps, _ = stats.Mean(weather.createTempList())
	return
}

func (weather *Weather) MedianTemp() (medianTemps float64) {
	medianTemps, _ = stats.Median(weather.createTempList())
	return
}

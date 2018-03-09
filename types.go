package main

import (
	"strconv"
	"strings"
	"time"
)


type Temp struct {
	TempF float64
	TempC float64
}

func (this *Temp) UnmarshalJSON(b []byte) error {
	this.TempC, _ = strconv.ParseFloat(string(b), 64)
	this.TempF = (this.TempC * 1.8 ) + 32

	return nil
}

type LatLong struct {
	Latitude  float64
	Longitude float64
}

func (this *LatLong) UnmarshalJSON(b []byte) error {
	unquoted := strings.Replace(string(b), "\"", "", -1)
	coords := strings.Split(unquoted, ",")
	latitude, _ := strconv.ParseFloat(coords[0], 64)
	longitude, _ := strconv.ParseFloat(coords[1], 64)
	this.Latitude = latitude
	this.Longitude = longitude

	return nil
}

type CityInfo struct {
	Title string

	Location_Type string
	Woeid         int
	Latt_Long     LatLong
}

type parent struct {
	Title         string
	Location_Type string
	Latt_Long     LatLong
	Woeid         int
}

type consolidatedWeather struct {
	Id                     int
	Applicable_Date        string
	Weather_State_Name     string
	Weather_State_Abbr     string
	Wind_Speed             float64
	Wind_Direction         float64
	Wind_Direction_Compass string
	Min_Temp               Temp
	Max_Temp               Temp
	The_Temp               Temp
	Air_Pressure           float64
	Humidity               float64
	Visibility             float64
	Predictability         int
}

type sources struct {
	Title string
	Url   string
}

type Weather struct {
	Title                string
	Location_Type        string
	Latt_Long            LatLong
	Time                 time.Time
	Sun_Rise             time.Time
	Sun_Set              time.Time
	Timezone             string
	Parent               parent
	Consolidated_Weather []consolidatedWeather
	Sources              []sources
}

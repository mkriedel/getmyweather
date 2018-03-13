package gmw

import (
	"time"
)

// Temp holds two float64 values representing temperatures in
// Celsius (TempC) and Fahrenheit (TempF)
type Temp struct {
	TempF float64
	TempC float64
}

// LatLong holds latitude and longitude values as float64s.
type LatLong struct {
	Latitude, Longitude float64
}

// Location type holds WOEID and other information about cities returned from the API.
type Location struct {
	Title         string
	Location_Type string
	Woeid         int
	Latt_Long     LatLong
}

type ConsolidatedWeather struct {
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

// Weather type for holding unmarshalled json data
type Weather struct {
	Title                string
	Location_Type        string
	Latt_Long            LatLong
	Time                 time.Time
	Sun_Rise             time.Time
	Sun_Set              time.Time
	Timezone             string
	Parent               Location
	Consolidated_Weather []ConsolidatedWeather
	Sources              []sources
}

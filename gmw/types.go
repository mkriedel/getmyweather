package gmw

import (
	"strconv"
	"strings"
	"time"
)

// Temp holds two float64 values representing temperatures in
// Celsius (TempC) and Fahrenheit (TempF)
type Temp struct {
	TempF float64
	TempC float64
}

// UnmarshalJSON interface loads the reported temperatures into the Temp
// type during a JSON unmarshaling.
func (this *Temp) UnmarshalJSON(b []byte) error {
	this.TempC, _ = strconv.ParseFloat(string(b), 64)
	this.TempF = (this.TempC * 1.8) + 32

	return nil
}

// LatLong holds latitude and longitude values as float64s.
type LatLong struct {
	Latitude, Longitude float64
}

// UnmarshalJSON takes a comma-seperated string ("Lat,Long") into a LatLong type.
func (this *LatLong) UnmarshalJSON(b []byte) error {
	unquoted := strings.Replace(string(b), "\"", "", -1)
	coords := strings.Split(unquoted, ",")
	if len(coords) != 2 {
		// If we're unable to split the Latt_Long value, set them to 0, 0
		this.Latitude = 0
		this.Longitude = 0
	} else {
		latitude, _ := strconv.ParseFloat(coords[0], 64)
		longitude, _ := strconv.ParseFloat(coords[1], 64)
		this.Latitude = latitude
		this.Longitude = longitude
	}

	return nil
}

// Location type holds WOEID and other information about cities returned from the API.
type Location struct {
	Title         string
	Location_Type string
	Woeid         int
	Latt_Long     LatLong
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
	Consolidated_Weather []consolidatedWeather
	Sources              []sources
}

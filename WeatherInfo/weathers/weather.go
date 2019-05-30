package weathers

import "fmt"

type Weather struct {
	Main        Main
	WeatherCond []WeatherCond `json:"weather"`
	Sys         Sys
	Wind        Wind
	CityName    string `json:"name"`
	Message     string
}

type WeatherCond struct {
	Description string `json:"description"`
}

type Main struct {
	Temperature float64 `json:"temp"`
	Humidity    int
}

type Wind struct {
	WindSpeed     float64 `json:"speed"`
	WindDirection float64 `json:"deg"`
	WindGust      float64 `json:"gust,omitempty"`
}

type Sys struct {
	Sunrise int64
	Sunset  int64
}

func (w Weather) GetTemperature() (temp float64) {
	return w.Main.Temperature
}

func (w Weather) GetCloudiness() (description string) {
	return w.WeatherCond[0].Description
}

func (w Weather) GetHumidity() (humidity int) {
	return w.Main.Humidity
}

func (w Weather) GetWind() (speed float64, gust float64, direction string) {

	direction, err := сonvertWind(w.Wind.WindDirection)

	if err != nil {
		fmt.Println("Error:", err)
	}

	return w.Wind.WindSpeed, w.Wind.WindGust, direction
}

func (w Weather) GetSun() (sunrise int64, sunset int64) {
	return w.Sys.Sunrise, w.Sys.Sunset
}

func (w Weather) GetInfo() (cityName string, message string) {
	return w.CityName, w.Message
}

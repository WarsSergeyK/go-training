package main

import (
	"fmt"
	"os"
	"weatherinfo/weathers"
)

func main() {

	var answer string

	fmt.Print("Укажите город: ")
	_, err := fmt.Scanln(&answer)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	var Meteorologist weathers.Meteorologist
	weather := Meteorologist.GetWeather(answer)

	var wf weathers.WeatherForecast
	fmt.Println(wf.FormatWeather(weather))

}

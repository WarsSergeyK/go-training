package weathers

import (
	"fmt"
	"os"
	"time"
)

type WeatherForecast struct{}

func (wf WeatherForecast) FormatWeather(w Weather) string {

	cityName, message := w.GetInfo()

	if message != "" {
		fmt.Println("Error:", message)
		os.Exit(1)
	}

	temp := w.GetTemperature()
	description := w.GetCloudiness()
	humidity := w.GetHumidity()
	speed, gust, direction := w.GetWind()
	sunrise, sunset := w.GetSun()

	var gustInfo string
	if gust > 0 {
		gustInfo = fmt.Sprintf(" с порывами до %.0fм/с", gust)
	}

	info := fmt.Sprintf("Сегодня в городе %s %s, температура воздуха %+.0f°С, ветер %s %.0fм/с%s. Влажность воздуха %d%%. Восход солнца %s, заход солнца %s.",
		cityName,
		description,
		temp,
		direction,
		speed,
		gustInfo,
		humidity,
		time.Unix(sunrise, 0).Format("15:04"),
		time.Unix(sunset, 0).Format("15:04"))

	return info
}

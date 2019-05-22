package weathers

import "errors"

func сonvertWind(d float64) (string, error) {

	switch {
	case d >= 337.5 && d <= 22.5:
		return "северный", nil //"N"

	case d > 22.5 && d < 67.5:
		return "северо-восточный", nil //"NE"

	case d >= 67.5 && d <= 112.5:
		return "восточный", nil //"E"

	case d > 112.5 && d < 157.5:
		return "юго-восточный", nil //"SE"

	case d >= 157.5 && d <= 202.5:
		return "южный", nil //"S"

	case d > 202.5 && d < 247.5:
		return "юго-западный", nil //"SW"

	case d >= 247.5 && d <= 292.5:
		return "западный", nil //"W"

	case d > 292.5 && d < 337.5:
		return "северо-западный", nil //"NW"

	default:
		return "неопределенный", errors.New("Cannot detect the wind direction")
	}
}

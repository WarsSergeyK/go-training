package weathers

import (
	"errors"
	"math"
)

const (
	circleDegrees = 360.0
	sectorsTotal  = 16
)

func сonvertWind(d float64) (string, error) {

	sector := math.Floor(d / circleDegrees * sectorsTotal)

	switch {
	case sector < 1 || sector > 15:
		return "северный", nil //"N"

	case sector <= 3:
		return "северо-восточный", nil //"NE"

	case sector <= 5:
		return "восточный", nil //"E"

	case sector <= 7:
		return "юго-восточный", nil //"SE"

	case sector <= 9:
		return "южный", nil //"S"

	case sector <= 11:
		return "юго-западный", nil //"SW"

	case sector <= 13:
		return "западный", nil //"W"

	case sector <= 15:
		return "северо-западный", nil //"NW"

	default:
		return "неопределенный", errors.New("Cannot parse the wind direction")
	}
}

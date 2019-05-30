package weathers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Meteorologist struct{}

func (m Meteorologist) GetWeather(city string) Weather {

	resp, err := http.Get(apiLink + "&q=" + city)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	var BodyBytes []byte

	if resp.StatusCode == http.StatusOK {
		BodyBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Not found")
		os.Exit(1)
	}

	var weather Weather
	err = json.Unmarshal(BodyBytes, &weather)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return weather
}
